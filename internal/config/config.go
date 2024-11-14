package config

import (
	"github.com/yogamandayu/ohmytp/pkg/db"
	"github.com/yogamandayu/ohmytp/pkg/redis"
	"github.com/yogamandayu/ohmytp/pkg/telegram"
	"github.com/yogamandayu/ohmytp/util"
)

// Config is a struct to hold all config.
type Config struct {
	REST                    *RESTConfig
	DB                      *DBConfig
	RedisAPI                *RedisConfig
	RedisWorkerNotification *RedisConfig
	TelegramBot             *TelegramBotConfig
}

// Option is an option for config.
type Option func(c *Config)

// NewConfig is a constructor.
func NewConfig() *Config {
	return &Config{}
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
		c.DB.Config = &db.Config{
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

// WithRedisAPIConfig is to set redis API config.
func WithRedisAPIConfig() Option {
	return func(c *Config) {
		if c.RedisAPI == nil {
			c.RedisAPI = &RedisConfig{}
		}
		c.RedisAPI.Config = &redis.Config{
			DB:       util.GetEnvAsInt("REDIS_API_DB", 0),
			Host:     util.GetEnv("REDIS_API_HOST", "localhost"),
			Port:     util.GetEnv("REDIS_API_PORT", "6379"),
			Password: util.GetEnv("REDIS_API_PASSWORD", "-"),
			PoolSize: util.GetEnvAsInt("REDIS_API_POOL_SIZE", 0),
		}
	}
}

// WithRedisWorkerNotificationConfig is to set redis worker config.
func WithRedisWorkerNotificationConfig() Option {
	return func(c *Config) {
		if c.RedisWorkerNotification == nil {
			c.RedisWorkerNotification = &RedisConfig{}
		}
		c.RedisWorkerNotification.Config = &redis.Config{
			DB:       util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_DB", 1),
			Host:     util.GetEnv("REDIS_WORKER_NOTIFICATION_HOST", "localhost"),
			Port:     util.GetEnv("REDIS_WORKER_NOTIFICATION_PORT", "6379"),
			Password: util.GetEnv("REDIS_WORKER_NOTIFICATION_PASSWORD", "-"),
			PoolSize: util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_POOL_SIZE", 0),
		}
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

// WithTelegramBotConfig is to set TelegramBot config.
func WithTelegramBotConfig() Option {
	return func(c *Config) {
		if c.TelegramBot == nil {
			c.TelegramBot = &TelegramBotConfig{}
		}
		c.TelegramBot.Config = &telegram.Config{
			Token:  util.GetEnv("TELEGRAM_BOT_TOKEN", "example-token"),
			ChatID: util.GetEnv("TELEGRAM_BOT_CHAT_ID", "example-id"),
		}
	}
}
