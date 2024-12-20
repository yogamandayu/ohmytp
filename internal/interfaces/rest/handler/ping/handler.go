package ping

import (
	"github.com/yogamandayu/ohmytp/internal/app"
)

// Handler is a struct to hold dependency.
type Handler struct {
	app *app.App
}

// NewHandler is a constructor.
func NewHandler(app *app.App) *Handler {
	return &Handler{
		app: app,
	}
}
