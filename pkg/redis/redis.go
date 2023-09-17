package redis

import (
	"go-gin-frame/config"
	"go-gin-frame/pkg/errors"
	"sync"

	"github.com/go-redis/redis/v7"
)

var (
	err    error
	client *redis.Client
	once   sync.Once
)

// 单例
func GetInstance() (*redis.Client, error) {
	once.Do(func() {
		cfg := config.Get().Redis
		client = redis.NewClient(&redis.Options{
			Addr:         cfg.Addr,
			Password:     cfg.Pass,
			DB:           cfg.Db,
			MaxRetries:   cfg.MaxRetries,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
		})

	})

	if err = client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}

	return client, nil
}
