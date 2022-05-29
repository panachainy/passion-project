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
	GetToday()
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

func (s *CovidServiceImp) GetToday() {
	client := resty.New()
	client.SetTimeout(30 * time.Second)

	url := "https://covid19.ddc.moph.go.th/api/Cases/today-cases-all"
	var result CovidClientResponse

	_, err := client.R().
		SetResult(&result).
		Get(url)
	if err != nil {
		logrus.Errorf("GetToday 1: %v", err.Error())
		// return "", err
	}

	logrus.Infoln("result", result)

	// return result.Result, nil
}
