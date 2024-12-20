package rest

import (
	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/config"
)

// Option is option to rest struct.
type Option func(r *REST)

// SetByConfig to set REST API with config.
func SetByConfig(config *config.Config) Option {
	return func(r *REST) {
		r.Port = config.REST.Port
	}
}

// WithApp to set app.
func WithApp(app *app.App) Option {
	return func(r *REST) {
		r.app = app
	}
}
