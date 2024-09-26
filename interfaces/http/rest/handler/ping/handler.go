package ping

import (
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	db *pgx.Conn
}

func NewHandler(db *pgx.Conn) *Handler {
	return &Handler{
		db,
	}
}
