//go:build wireinject
// +build wireinject

//go:generate wire
package build

import (
	"covid-19-api/cmd/config"

	"github.com/google/wire"
)

func Wire(envFile string) (*ApplicationImp, error) {
	wire.Build(
		// TODO: wait to implement service to use cache
		// cache.ProviderSet,
		ProviderApp,
		config.ProviderSet,
	)

	return &ApplicationImp{}, nil
}
