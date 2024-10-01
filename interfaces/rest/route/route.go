package route

import (
	"net/http"

	"github.com/yogamandayu/ohmytp/app"
)

type Router struct {
	app *app.App
}

func NewRouter(app *app.App) *Router {
	return &Router{
		app,
	}
}

func (r *Router) Handler() http.Handler {
	mux := http.NewServeMux()
	HealthRoute(mux, r.app)
	OTPRoute(mux, r.app)

	return mux
}
