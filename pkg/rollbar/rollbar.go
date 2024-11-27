package rollbar

import (
	"github.com/rollbar/rollbar-go"
)

type Config struct {
	Environment string
	Token       string
	CodeVersion string
	Host        string
	ServerRoot  string
}

func NewRollbar(config *Config) *rollbar.Client {
	return rollbar.New(config.Token, config.Environment, config.CodeVersion, config.Host, config.ServerRoot)
}
