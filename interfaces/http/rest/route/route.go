package route

import "net/http"

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handler() http.Handler {

	return HealthRoute()
}
