package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Config is a redis config.
type Config struct {
	Password string
	Host     string
	Port     string
	DB       int
	PoolSize int
}

// NewConnection is to set new redis connection.
func NewConnection(config *Config) (*redis.Client, error) {
	if config == nil {
		return nil, errors.New("redis.error.missing_config")
	}

	rdb := redis.NewClient(&redis.Options{
		DB:       config.DB,
		Password: config.Password,
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		PoolSize: func() int {
			return config.PoolSize
		}(),
	})

	redisStatus := rdb.Ping(context.Background())
	if redisStatus.Err() != nil {
		return nil, redisStatus.Err()
	}

	return rdb, nil
}
