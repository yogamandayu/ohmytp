package app

import (
	"log/slog"

	"github.com/yogamandayu/ohmytp/internal/config"

	"github.com/yogamandayu/ohmytp/internal/storage/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
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

func WithConfig(config *config.Config) Option {
	return func(a *App) {
		a.Config = config
	}
}
