package config

import (
	"github.com/yogamandayu/ohmytp/util"
)

type Config struct {
	REST  *RESTConfig
	DB    *DBConfig
	Redis *RedisConfig
}

type Option func(c *Config)

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) With(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func WithDBConfig() Option {
	return func(c *Config) {
		c.DB = &DBConfig{
			Driver:                      util.GetEnv("DB_DRIVER", "mysql"),
			Host:                        util.GetEnv("DB_HOST", "localhost"),
			Port:                        util.GetEnv("DB_PORT", "3306"),
			Username:                    util.GetEnv("DB_USER", "root"),
			Password:                    util.GetEnv("DB_PASSWORD", "-"),
			Database:                    util.GetEnv("DB_NAME", "ohmytp"),
			TimeZone:                    util.GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
			Log:                         util.GetEnvAsBool("ENABLE_LOGGER", true),
			DisableForeignKeyConstraint: util.GetEnvAsBool("DISABLE_FOREIGN_KEY_CONSTRAINT", false),
		}
	}
}

func WithRedisConfig() Option {
	return func(c *Config) {
		c.Redis = &RedisConfig{
			Host:     util.GetEnv("REDIS_HOST", "localhost"),
			Port:     util.GetEnv("REDIS_PORT", "6379"),
			Password: util.GetEnv("REDIS_PASSWORD", "-"),
		}
	}
}

func WithRESTConfig() Option {
	return func(c *Config) {
		c.REST = &RESTConfig{
			Port: util.GetEnv("APP_PORT", "8080"),
		}
	}
}
