package app

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// App is a struct to hold all dependecy that used in the otp service.
type App struct {
	DB    *pgxpool.Pool
	Redis *redis.Client
	Log   *slog.Logger
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
