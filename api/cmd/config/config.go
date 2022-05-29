package config

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configOnce sync.Once
	c          = Configuration{
		APP_ENV:   "development",
		APP_PORT:  9090,
		LOG_LEVEL: "debug",

		DATABASE_TEST: false,

		REDIS_TLS: false,
	}
)

type Configuration struct {
	APP_ENV        string `mapstructure:"APP_ENV"`
	APP_PORT       int    `mapstructure:"APP_PORT" validate:"required"`
	APP_DEBUG      bool   `mapstructure:"APP_DEBUG"`
	APP_PREFIX     string `mapstructure:"APP_PREFIX"`
	APP_PUBLIC_URL string `mapstructure:"APP_PUBLIC_URL"`

	LOG_LEVEL string `mapstructure:"LOG_LEVEL"`

	DATABASE_DSN  string `mapstructure:"DATABASE_DSN" validate:"required"`
	DATABASE_TEST bool   `mapstructure:"DATABASE_TEST"`

	REDIS_HOST     string `mapstructure:"REDIS_HOST"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_PORT     string `mapstructure:"REDIS_PORT"`
	REDIS_DB       int    `mapstructure:"REDIS_DB"`
	REDIS_TLS      bool   `mapstructure:"REDIS_TLS"`
}

type ErrorResponse struct {
	FailedField string `json:"failedFields"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ProviderConfig(filePath string) *Configuration {
	configOnce.Do(func() {
		if filePath != "" {
			viper.SetConfigFile(filePath)

			err := viper.ReadInConfig()
			if err != nil {
				panic(err)
			}
		}

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		bindEnvs(c)

		err := viper.Unmarshal(&c)
		if err != nil {
			logrus.Errorf("unable to decode into struct, %v", err)
			panic("[CONFIG] unable to decode into struct")
		}

		e := Validate(c)
		if e != nil {
			for _, s := range e {
				logrus.Errorf("[CONFIG] %s: %s\n", s.FailedField, s.Tag)
			}
			panic("[CONFIG] invalid configuration")
		}
	})

	return &c
}

func PrintConfig() {
	logrus.Debugf("[CONFIG]: %#v\n", c)
}

func Validate(c Configuration) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			err := viper.BindEnv(strings.Join(append(parts, tv), "."))
			if err != nil {
				fmt.Printf("can't bind config from ENV, %v\n", err)
			}
		}
	}
}
