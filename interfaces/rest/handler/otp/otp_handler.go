package otp

import (
	"context"
	"net/http"
	"time"

	"github.com/yogamandayu/ohmytp/interfaces/rest/response"

	"encoding/json"

	"github.com/yogamandayu/ohmytp/consts"
	"github.com/yogamandayu/ohmytp/requester"
	"github.com/yogamandayu/ohmytp/workflow/otp"
)

// Request is request otp request handler.
func (h *Handler) Request(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)

	var payload RequestOtpRequestContract
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		h.app.Log.Warn(err.Error())
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusBadRequest).AsJSON(w)
		return
	}

	rq := requester.NewRequester().SetMetadataFromREST(r)
	workflow := otp.NewRequestOtpWorkflow(rq, h.app)

	otpEntity := payload.TransformToOtpEntity()
	workflow.SetOtp(&otpEntity).SetOtpExpiration(time.Duration(payload.Expiration) * time.Second).SetOtpLength(payload.Length)
	switch otpEntity.RouteType {
	case consts.EmailRouteType.ToString():
		_ = workflow.WithRouteEmail(payload.RouteValue)
	case consts.SMSRouteType.ToString():
		_ = workflow.WithRouteSMS(payload.RouteValue)
	}
	expiredAt, err := workflow.Request(ctx)
	if err != nil {
		h.app.Log.Error(err.Error())
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusBadRequest).AsJSON(w)
		return
	}
	response.NewHTTPSuccessResponse(RequestOtpResponseContract{
		ExpiredAt: expiredAt.Format(time.RFC3339),
	}, "Success").WithStatusCode(http.StatusCreated).AsJSON(w)
}