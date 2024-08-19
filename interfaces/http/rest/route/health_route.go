package route

import (
	"net/http"

	"github.com/yogamandayu/ohmytp/interfaces/http/rest/handler/ping"
)

func HealthRoute() http.Handler {
	mux := http.NewServeMux()

	pingHandler := ping.NewHandler()

	mux.HandleFunc("/ping", pingHandler.Ping)

	return mux
}
