package app

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/yogamandayu/ohmytp/storage/repository"
)

type Option func(*App)

func WithDB(pgxConn *pgxpool.Pool) Option {
	return func(a *App) {
		a.DB = pgxConn
	}
}

func WithDBRepository(pgxConn *pgxpool.Pool) Option {
	return func(a *App) {
		a.DBRepository = repository.New(pgxConn)
	}
}

func WithRedis(redisConn *redis.Client) Option {
	return func(a *App) {
		a.Redis = redisConn
	}
}

func WithSlog(slog *slog.Logger) Option {
	return func(a *App) {
		a.Log = slog
	}
}
