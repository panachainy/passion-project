//go:generate mockgen -source=service.go -destination=mock/mock_service.go -package=mock
package covid

import (
	"sync"

	"covid-19-api/cmd/config"
)

var (
	sOnce           sync.Once
	serviceInstance *CovidServiceImp
)

type CovidService interface{}

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
	// https://covid19.ddc.moph.go.th/api/Cases/today-cases-all
}
