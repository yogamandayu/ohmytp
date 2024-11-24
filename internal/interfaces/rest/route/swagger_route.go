package route

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/yogamandayu/ohmytp/internal/app"
)

// SwaggerRoute is a route for swagger API doc.
func SwaggerRoute(mux *http.ServeMux, app *app.App) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
