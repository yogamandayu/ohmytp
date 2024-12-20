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

// App is a struct to hold all dependency that used in the otp service.
type App struct {
	DB                      *pgxpool.Pool
	RedisAPI                *redis.Client
	RedisWorkerNotification *redis.Client
	DBRepository            *repository.Queries
	Log                     *slog.Logger
	Rollbar                 *rollbar.Client
	Minio                   *minio.Client
	Config                  *config.Config
}

// NewApp is a constructor.
func NewApp() *App {
	return &App{}
}

// WithOptions is to set options to app.
func (app *App) WithOptions(opts ...Option) *App {
	for _, opt := range opts {
		opt(app)
	}
	return app
}
