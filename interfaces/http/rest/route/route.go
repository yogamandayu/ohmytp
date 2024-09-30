package route

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handler(db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()
	HealthRoute(mux, db)

	return mux
}
