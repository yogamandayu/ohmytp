package route

import (
	"net/http"

	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/handler/ping"
)

// HealthRoute is a health route to monitor service health.
func HealthRoute(mux *http.ServeMux, app *app.App) {
	pingHandler := ping.NewHandler(app.DB, app.Redis)
	mux.HandleFunc("/ping", pingHandler.Ping)
}
