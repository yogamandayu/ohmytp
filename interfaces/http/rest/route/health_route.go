package route

import (
	"net/http"

	"github.com/yogamandayu/ohmytp/interfaces/http/rest/handler/ping"
)

func HealthRoute(mux *http.ServeMux) http.Handler {
	pingHandler := ping.NewHandler()
	mux.HandleFunc("/ping", pingHandler.Ping)

	return mux
}
