package rest

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/yogamandayu/ohmytp/config"
)

func WithConfig(config *config.Config) Option {
	return func(r *REST) {
		r.config = config
	}
}

func WithDB(db *sql.DB) Option {
	return func(r *REST) {
		r.db = db
	}
}

func WithRedis(redis *redis.Client) Option {
	return func(r *REST) {
		r.redis = redis
	}
}
