package cache

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProviderRedis,
	wire.Bind(new(CacheInterface), new(*RedisImpl)),
)
