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
		if c.DB == nil {
			c.DB = &DBConfig{}
		}
		c.DB.Driver = util.GetEnv("DB_DRIVER", "mysql")
		c.DB.Host = util.GetEnv("DB_HOST", "localhost")
		c.DB.Port = util.GetEnv("DB_PORT", "3306")
		c.DB.Username = util.GetEnv("DB_USER", "root")
		c.DB.Password = util.GetEnv("DB_PASSWORD", "-")
		c.DB.Database = util.GetEnv("DB_NAME", "ohmytp")
		c.DB.TimeZone = util.GetEnv("APP_TIMEZONE", "Asia/Jakarta")
		c.DB.Log = util.GetEnvAsBool("DB_LOGGER", false)
	}
}

// WithRedisConfig is to set redis config.
func WithRedisConfig() Option {
	return func(c *Config) {
		if c.Redis == nil {
			c.Redis = &RedisConfig{}
		}
		c.Redis.DB = util.GetEnvAsInt("REDIS_DB", 0)
		c.Redis.Host = util.GetEnv("REDIS_HOST", "localhost")
		c.Redis.Port = util.GetEnv("REDIS_PORT", "6379")
		c.Redis.Password = util.GetEnv("REDIS_PASSWORD", "-")
		c.Redis.PoolSize = util.GetEnvAsInt("REDIS_POOL_SIZE", 0)
	}
}

// WithRESTConfig is to set REST API config.
func WithRESTConfig() Option {
	return func(c *Config) {
		if c.REST == nil {
			c.REST = &RESTConfig{}
		}
		c.REST = &RESTConfig{
			Port: util.GetEnv("APP_PORT", "8080"),
		}
	}
}
