package otp

import (
	"github.com/yogamandayu/ohmytp/app"
)

// Handler is a struct to hold dependency.
type Handler struct {
	app *app.App
}

// NewHandler is a constructor.
func NewHandler(app *app.App) *Handler {
	return &Handler{
		app,
	}
}
