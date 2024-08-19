package config

import "github.com/yogamandayu/ohmytp/util"

type Config struct {
	REST *RESTConfig
	DB   *DBConfig
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
			User:                        util.GetEnv("DB_USER", "root"),
			Name:                        util.GetEnv("DB_NAME", "ohmytp"),
			Password:                    util.GetEnv("DB_PASSWORD", "-"),
			TimeZone:                    util.GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
			Log:                         util.GetEnvAsBool("ENABLE_LOGGER", true),
			DisableForeignKeyConstraint: util.GetEnvAsBool("DISABLE_FOREIGN_KEY_CONSTRAINT", false),
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
