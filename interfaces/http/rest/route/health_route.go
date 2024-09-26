package route

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/yogamandayu/ohmytp/interfaces/http/rest/handler/ping"
)

func HealthRoute(mux *http.ServeMux, db *pgx.Conn) http.Handler {
	pingHandler := ping.NewHandler(db)
	mux.HandleFunc("/ping", pingHandler.Ping)

	return mux
}
