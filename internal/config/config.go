package config

import (
	"github.com/yogamandayu/ohmytp/util"
)

// Config is a struct to hold all config.
type Config struct {
	REST  *RESTConfig
	DB    *DBConfig
	Redis *RedisConfig
}

// Option is an option for config.
type Option func(c *Config)

// NewConfig is a constructor.
func NewConfig() Config {
	return Config{}
}

// WithOptions is to set option to config.
func (c *Config) WithOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

// WithDBConfig is to set db config.
func WithDBConfig() Option {
	return func(c *Config) {
		c.DB = &DBConfig{
			Driver:   util.GetEnv("DB_DRIVER", "mysql"),
			Host:     util.GetEnv("DB_HOST", "localhost"),
			Port:     util.GetEnv("DB_PORT", "3306"),
			Username: util.GetEnv("DB_USER", "root"),
			Password: util.GetEnv("DB_PASSWORD", "-"),
			Database: util.GetEnv("DB_NAME", "ohmytp"),
			TimeZone: util.GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
			Log:      util.GetEnvAsBool("DB_LOGGER", false),
		}
	}
}

// WithRedisConfig is to set redis config.
func WithRedisConfig() Option {
	return func(c *Config) {
		c.Redis = &RedisConfig{
			DB:       util.GetEnvAsInt("REDIS_DB", 0),
			Host:     util.GetEnv("REDIS_HOST", "localhost"),
			Port:     util.GetEnv("REDIS_PORT", "6379"),
			Password: util.GetEnv("REDIS_PASSWORD", "-"),
			PoolSize: util.GetEnvAsInt("REDIS_POOL_SIZE", 0),
		}
	}
}

// WithRESTConfig is to set REST API config.
func WithRESTConfig() Option {
	return func(c *Config) {
		c.REST = &RESTConfig{
			Port: util.GetEnv("APP_PORT", "8080"),
		}
	}
}
