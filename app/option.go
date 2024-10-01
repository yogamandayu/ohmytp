package app

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AppOption func(*App)

func WithDB(pgxConn *pgxpool.Pool) AppOption {
	return func(a *App) {
		a.DB = pgxConn
	}
}

func WithRedis(redisConn *redis.Client) AppOption {
	return func(a *App) {
		a.Redis = redisConn
	}
}

func WithSlog(slog *slog.Logger) AppOption {
	return func(a *App) {
		a.Log = slog
	}
}
