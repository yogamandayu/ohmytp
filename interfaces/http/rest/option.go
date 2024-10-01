package rest

import (
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/config"
)

type Option func(r *REST)

func WithConfig(config *config.Config) Option {
	return func(r *REST) {
		r.Port = config.REST.Port
	}
}

func WithApp(app *app.App) Option {
	return func(r *REST) {
		r.app = app
	}
}
