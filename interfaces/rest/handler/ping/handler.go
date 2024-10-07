package ping

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// Handler is a struct to hold dependency.
type Handler struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

// NewHandler is a constructor.
func NewHandler(db *pgxpool.Pool, redis *redis.Client) *Handler {
	return &Handler{
		db:    db,
		redis: redis,
	}
}
