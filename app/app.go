package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	DB    *pgxpool.Pool
	Redis *redis.Client
}

func NewApp() *App {
	return &App{}
}

func (app *App) WithOptions(opts ...AppOption) *App {
	for _, opt := range opts {
		opt(app)
	}
	return app
}
