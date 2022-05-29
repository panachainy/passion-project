package build

import (
	"sync"

	"covid-19-api/cmd/config"
	"covid-19-api/internal/covid"
)

var (
	appOnce sync.Once
	app     = &ApplicationImp{}
)

type ApplicationImp struct {
	Config       *config.Configuration
	CovidHandler covid.CovidHandler
}

func ProviderApp(c *config.Configuration, ch covid.CovidHandler) *ApplicationImp {
	appOnce.Do(func() {
		app = &ApplicationImp{
			Config:       c,
			CovidHandler: ch,
		}
	})

	return app
}
