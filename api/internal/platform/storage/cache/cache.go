//go:generate mockgen -source=cache.go -destination=mock/mock_cache.go -package=mock

package cache

import "time"

type CacheInterface interface {
	Get(key string, dest interface{}) error
	Set(key string, value interface{}, expiration time.Duration) error
	DeleteWithPrefix(prefix string) error
	Delete(key string) error
}
