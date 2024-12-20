package app

import (
	"log/slog"

	"github.com/minio/minio-go/v7"

	"github.com/rollbar/rollbar-go"

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

func WithRedisAPI(redisConn *redis.Client) Option {
	return func(a *App) {
		a.RedisAPI = redisConn
	}
}

func WithRedisWorkerNotification(redisConn *redis.Client) Option {
	return func(a *App) {
		a.RedisWorkerNotification = redisConn
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

func WithRollbar(rollbar *rollbar.Client) Option {
	return func(a *App) {
		a.Rollbar = rollbar
	}
}

func WithMinio(minio *minio.Client) Option {
	return func(a *App) {
		a.Minio = minio
	}
}
