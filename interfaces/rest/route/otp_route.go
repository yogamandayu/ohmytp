package route

import (
	"net/http"

	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/interfaces/rest/handler/otp"
)

// OTPRoute is otp route.
func OTPRoute(mux *http.ServeMux, app *app.App) {
	otpHandler := otp.NewHandler(app)

	groupV1 := "/api/v1"
	mux.HandleFunc(Group("POST", groupV1, "otp/request"), otpHandler.Request)
	mux.HandleFunc(Group("POST", groupV1, "otp/confirm"), otpHandler.Confirm)
}
