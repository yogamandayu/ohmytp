package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config is a redis config.
type Config struct {
	Password        string
	Host            string
	Port            string
	DB              int
	PoolSize        int
	MinIdleConns    int
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	MaxConnIdleTime time.Duration
	MaxConnLifeTime time.Duration
}

// NewConnection is to set new redis connection.
func NewConnection(config *Config) (*redis.Client, error) {
	if config == nil {
		return nil, errors.New("redis.error.missing_config")
	}

	rdb := redis.NewClient(&redis.Options{
		DB:              config.DB,
		Password:        config.Password,
		Addr:            fmt.Sprintf("%s:%s", config.Host, config.Port),
		PoolSize:        config.PoolSize,
		MinIdleConns:    config.MinIdleConns,
		DialTimeout:     config.DialTimeout,
		ReadTimeout:     config.ReadTimeout,
		WriteTimeout:    config.WriteTimeout,
		ConnMaxIdleTime: config.MaxConnIdleTime,
		ConnMaxLifetime: config.MaxConnLifeTime,
	})

	redisStatus := rdb.Ping(context.Background())
	if redisStatus.Err() != nil {
		return nil, redisStatus.Err()
	}

	return rdb, nil
}
