package redis

import (
	"context"
	"fmt"

	"github.com/yogamandayu/ohmytp/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewConnection(config config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		DB:       config.Redis.DB,
		Password: config.Redis.Password,
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		PoolSize: func() int {
			return config.Redis.PoolSize
		}(),
	})

	redisStatus := rdb.Ping(context.Background())
	if redisStatus.Err() != nil {
		return nil, redisStatus.Err()
	}

	return rdb, nil
}
