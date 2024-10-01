package route

import (
	"net/http"

	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/interfaces/http/rest/handler/ping"
)

func HealthRoute(mux *http.ServeMux, app *app.App) http.Handler {
	pingHandler := ping.NewHandler(app.DB, app.Redis)
	mux.HandleFunc("/ping", pingHandler.Ping)

	return mux
}
