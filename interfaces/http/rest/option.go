package rest

import "github.com/yogamandayu/ohmytp/config"

func WithConfig(config *config.Config) Option {
	return func(r *REST) {
		r.Config = config
	}
}
