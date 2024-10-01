package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yogamandayu/ohmytp/config"
)

func NewConnection(config config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		DB:       config.Redis.DB,
		Password: config.Redis.Password,
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
	})

	redisStatus := rdb.Ping(context.Background())
	if redisStatus.Err() != nil {
		return nil, redisStatus.Err()
	}

	return rdb, nil
}
