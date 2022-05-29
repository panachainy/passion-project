package covid

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProviderService,
	wire.Bind(new(CovidService), new(*CovidServiceImp)),

	ProviderHandler,
	wire.Bind(new(CovidHandler), new(*CovidHandlerImp)),
)
