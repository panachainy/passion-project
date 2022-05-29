//go:build wireinject
// +build wireinject

//go:generate wire
package build

import (
	"covid-19-api/cmd/config"

	"covid-19-api/internal/platform/storage/cache"

	"github.com/google/wire"
)

func Wire(envFile string) (*ApplicationImp, error) {
	wire.Build(
		cache.ProviderSet,
		ProviderApp,
		config.ProviderSet,
	)

	return &ApplicationImp{}, nil
}
