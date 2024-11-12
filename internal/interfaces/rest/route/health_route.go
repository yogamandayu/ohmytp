package route

import (
	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/handler/ping"
	"net/http"
)

// HealthRoute is a health route to monitor service health.
func HealthRoute(mux *http.ServeMux, app *app.App) {
	pingHandler := ping.NewHandler(app.DB, app.Redis)
	mux.HandleFunc("/ping", pingHandler.Ping)
}
