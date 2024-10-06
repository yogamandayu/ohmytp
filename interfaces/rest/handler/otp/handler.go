package otp

import (
	"github.com/yogamandayu/ohmytp/app"
)

type Handler struct {
	app *app.App
}

func NewHandler(app *app.App) *Handler {
	return &Handler{
		app,
	}
}
