package build

import (
	"sync"

	"covid-19-api/cmd/config"
)

var (
	appOnce sync.Once
	app     = &ApplicationImp{}
)

type ApplicationImp struct {
	Config *config.Configuration
}

func ProviderApp(c *config.Configuration) *ApplicationImp {
	appOnce.Do(func() {
		app = &ApplicationImp{
			Config: c,
		}
	})

	return app
}
