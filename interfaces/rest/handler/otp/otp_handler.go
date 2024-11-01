package otp

import (
	"context"
	"log"
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
		log.Fatal(err)
	}

	rq := requester.NewRequester().SetMetadataFromREST(r)
	workflow := otp.NewRequestOtpWorkflow(rq, h.app)

	otpEntity := payload.TransformToOtpEntity()
	workflow.SetOtp(&otpEntity)
	switch otpEntity.RouteType {
	case consts.EmailRouteType.ToString():
		_ = workflow.WithRouteEmail(payload.RouteValue)
	case consts.SMSRouteType.ToString():
		_ = workflow.WithRouteSMS(payload.RouteValue)
	}
	err = workflow.Request(ctx)
	if err != nil {
		h.app.Log.Error("error : %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	response.NewHTTPSuccessResponse(w, http.StatusCreated, nil, "Success")
}
