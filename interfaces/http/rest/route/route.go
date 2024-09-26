package route

import (
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handler(db *pgx.Conn) http.Handler {
	mux := http.NewServeMux()
	HealthRoute(mux, db)

	return mux
}
