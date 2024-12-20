package config

import (
	"time"

	"github.com/yogamandayu/ohmytp/pkg/minio"

	"github.com/yogamandayu/ohmytp/pkg/rollbar"

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
	Rollbar                 *RollbarConfig
	Minio                   *MinioConfig
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
			Host:              util.GetEnv("DB_HOST", "localhost"),
			Port:              util.GetEnv("DB_PORT", "3306"),
			Username:          util.GetEnv("DB_USER", "root"),
			Password:          util.GetEnv("DB_PASSWORD", "-"),
			Database:          util.GetEnv("DB_NAME", "ohmytp"),
			TimeZone:          util.GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
			Log:               util.GetEnvAsBool("DB_LOGGER", false),
			MaxConns:          util.GetEnvAsInt("DB_MAX_CONNS", 0),
			MinConns:          util.GetEnvAsInt("DB_MIN_CONNS", 0),
			MaxConnIdleTime:   time.Duration(util.GetEnvAsInt("DB_MAX_CONN_IDLE_TIME", 0)) * time.Second,
			MaxConnLifeTime:   time.Duration(util.GetEnvAsInt("DB_MAX_CONN_LIFE_TIME", 0)) * time.Second,
			HealthCheckPeriod: time.Duration(util.GetEnvAsInt("DB_HEALTH_CHECK_PERIOD", 1)) * time.Second,
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
			Password:        util.GetEnv("REDIS_API_PASSWORD", "-"),
			Host:            util.GetEnv("REDIS_API_HOST", "localhost"),
			Port:            util.GetEnv("REDIS_API_PORT", "6379"),
			DB:              util.GetEnvAsInt("REDIS_API_DB", 0),
			PoolSize:        util.GetEnvAsInt("REDIS_API_POOL_SIZE", 0),
			MinIdleConns:    util.GetEnvAsInt("REDIS_API_MIN_IDLE_CONNS", 0),
			DialTimeout:     time.Duration(util.GetEnvAsInt("REDIS_API_DIAL_TIMEOUT", 0)) * time.Second,
			ReadTimeout:     time.Duration(util.GetEnvAsInt("REDIS_API_READ_TIMEOUT", 0)) * time.Second,
			WriteTimeout:    time.Duration(util.GetEnvAsInt("REDIS_API_WRITE_TIMEOUT", 0)) * time.Second,
			MaxConnIdleTime: time.Duration(util.GetEnvAsInt("REDIS_API_MAX_CONN_IDLE_TIME", 0)) * time.Second,
			MaxConnLifeTime: time.Duration(util.GetEnvAsInt("REDIS_API_MAX_CONN_LIFE_TIME", 0)) * time.Second,
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
			DB:              util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_DB", 1),
			Host:            util.GetEnv("REDIS_WORKER_NOTIFICATION_HOST", "localhost"),
			Port:            util.GetEnv("REDIS_WORKER_NOTIFICATION_PORT", "6379"),
			Password:        util.GetEnv("REDIS_WORKER_NOTIFICATION_PASSWORD", "-"),
			PoolSize:        util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_POOL_SIZE", 0),
			MinIdleConns:    util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_MIN_IDLE_CONNS", 0),
			DialTimeout:     time.Duration(util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_DIAL_TIMEOUT", 0)) * time.Second,
			ReadTimeout:     time.Duration(util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_READ_TIMEOUT", 0)) * time.Second,
			WriteTimeout:    time.Duration(util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_WRITE_TIMEOUT", 0)) * time.Second,
			MaxConnIdleTime: time.Duration(util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_MAX_CONN_IDLE_TIME", 0)) * time.Second,
			MaxConnLifeTime: time.Duration(util.GetEnvAsInt("REDIS_WORKER_NOTIFICATION_MAX_CONN_LIFE_TIME", 0)) * time.Second,
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

// WithRollbarConfig is to set Rollbar config.
func WithRollbarConfig() Option {
	return func(c *Config) {
		if c.Rollbar == nil {
			c.Rollbar = &RollbarConfig{}
		}
		c.Rollbar.Config = &rollbar.Config{
			Token:       util.GetEnv("ROLLBAR_TOKEN", "example-token"),
			Environment: util.GetEnv("ROLLBAR_ENVIRONMENT", "development"),
		}
	}
}

// WithMinioConfig is to set minio config.
func WithMinioConfig() Option {
	return func(c *Config) {
		if c.Minio == nil {
			c.Minio = &MinioConfig{}
		}
		c.Minio.Config = &minio.Config{
			Host:            util.GetEnv("MINIO_HOST", "localhost"),
			AccessKeyID:     util.GetEnv("MINIO_ACCESS_KEY", "example-access-key"),
			SecretAccessKey: util.GetEnv("MINIO_SECRET_KEY", "example-secret-key"),
		}
	}
}
