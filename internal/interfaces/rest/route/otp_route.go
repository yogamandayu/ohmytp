package route

import (
	"net/http"
	"time"

	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/handler/otp"
	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/middleware"
)

// OTPRoute is route for all otp handler.
func OTPRoute(mux *http.ServeMux, app *app.App) {
	otpHandler := otp.NewHandler(app)

	mux.HandleFunc("POST /api/v1/otp/request", middleware.RateLimit(app, "otp:v1:request_otp").
		WithFixedWindow(1000, 1*time.Minute).
		LimitByIPAddress().
		Apply(
			http.HandlerFunc(otpHandler.Request),
		).ServeHTTP)
	mux.HandleFunc("POST /api/v1/otp/confirm", middleware.RateLimit(app, "otp:v1:confirm_otp").
		WithFixedWindow(1000, 1*time.Minute).
		LimitByIPAddress().
		Apply(
			http.HandlerFunc(otpHandler.Confirm),
		).ServeHTTP)
}