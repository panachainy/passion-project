//go:generate mockgen -source=service.go -destination=mock/mock_service.go -package=mock
package covid

import (
	"sync"
	"time"

	"covid-19-api/cmd/config"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

var (
	sOnce           sync.Once
	serviceInstance *CovidServiceImp
)

type CovidService interface {
	GetToday() (CovidClientResponse, error)
}

type CovidServiceImp struct {
	Config *config.Configuration
}

func ProviderService(c *config.Configuration) *CovidServiceImp {
	sOnce.Do(func() {
		serviceInstance = &CovidServiceImp{
			Config: c,
		}
	})

	return serviceInstance
}

func (s *CovidServiceImp) GetToday() (CovidClientResponse, error) {
	client := resty.New()
	client.SetTimeout(30 * time.Second)

	url := "https://covid19.ddc.moph.go.th/api/Cases/today-cases-all"
	var result CovidClientResponse

	if _, err := client.R().
		SetResult(&result).
		Get(url); err != nil {
		logrus.Errorf("GetToday 1: %v", err.Error())
		return nil, err
	}

	return result, nil
}
