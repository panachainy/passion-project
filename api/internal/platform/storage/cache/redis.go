package cache

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"sync"
	"time"

	"covid-19-api/cmd/config"

	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

var (
	redisOnce     sync.Once
	redisInstance *RedisImpl
)

type RedisImpl struct {
	rdClient *redis.Client
	ctx      context.Context
}

func ProviderRedis(conf *config.Configuration) *RedisImpl {
	redisOnce.Do(func() {
		redisRepo := RedisImpl{ctx: context.Background()}

		host := conf.REDIS_HOST
		pass := conf.REDIS_PASSWORD
		port := conf.REDIS_PORT
		dbName := conf.REDIS_DB

		addr := host + ":" + port

		if conf.REDIS_TLS {
			redisRepo.rdClient = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: pass,
				DB:       dbName,
				TLSConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
					ServerName: host,
				},
			})
		} else {
			redisRepo.rdClient = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: pass,
				DB:       dbName,
			})
		}

		// Check cache set or not
		pong, err := redisRepo.rdClient.Ping(redisRepo.ctx).Result()
		if err != nil {
			logrus.Warnf("[REDIS] Health error: %v", err)
		} else {
			logrus.Infof("[REDIS] Health result: %v", pong)
		}

		redisInstance = &redisRepo
	})

	return redisInstance
}

func (repo *RedisImpl) Set(key string, value interface{}, expiration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		logrus.Warnln("[REDIS] Can't Marshal: ", key, " to redis", err)
		return err
	}

	err = repo.rdClient.Set(repo.ctx, key, v, expiration).Err()
	if err != nil {
		logrus.Warnln("[REDIS] Can't set Key: ", key, " to redis", err)
	}

	return nil
}

func (repo *RedisImpl) Get(key string, dest interface{}) error {
	p, err := repo.rdClient.Get(repo.ctx, key).Result()
	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(p), dest); err != nil {
		return err
	}

	return nil
}

func (repo *RedisImpl) DeleteWithPrefix(prefix string) error {
	iter := repo.rdClient.Scan(repo.ctx, 0, prefix+"*", 0).Iterator()
	for iter.Next(repo.ctx) {
		err := repo.Delete(iter.Val())
		if err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}

func (repo *RedisImpl) Delete(key string) error {
	if err := repo.rdClient.Del(repo.ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
