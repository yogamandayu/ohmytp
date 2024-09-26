package rest

import (
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5"
	"github.com/yogamandayu/ohmytp/config"
)

func WithConfig(config *config.Config) Option {
	return func(r *REST) {
		r.config = config
	}
}

func WithDB(db *pgx.Conn) Option {
	return func(r *REST) {
		r.db = db
	}
}

func WithRedis(redis *redis.Client) Option {
	return func(r *REST) {
		r.redis = redis
	}
}
