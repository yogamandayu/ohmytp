package config

type Config struct {
	REST *RESTConfig
}

func NewConfig() *Config {
	return &Config{}
}
