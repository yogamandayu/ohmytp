package route

import (
	"net/http"

	"github.com/yogamandayu/ohmytp/app"
)

// Router is a struct to hold all dependency to provide route.
type Router struct {
	app *app.App
}

// NewRouter is a contructor.
func NewRouter(app *app.App) *Router {
	return &Router{
		app,
	}
}

// Handler is to get all route handler.
func (r *Router) Handler() http.Handler {
	mux := http.NewServeMux()
	HealthRoute(mux, r.app)
	OTPRoute(mux, r.app)

	return mux
}
