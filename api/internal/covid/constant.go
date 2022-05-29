package covid

import "time"

const (
	CACHE_PREFIX = "passion:"

	TODAY_CACHE_KEY = CACHE_PREFIX + "today:"
	CACHE_TIME      = time.Duration(60) * time.Second
)
