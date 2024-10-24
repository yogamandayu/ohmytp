package otp

import (
	"context"
	"log"
	"net/http"
	"time"

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
	requester := requester.NewRequester().SetMetadataFromREST(r)
	workflow := otp.NewRequestOtpWorkflow(requester, h.app)

	data := RequestOtpResponseContract{
		Message: "OK",
	}
	otp := payload.TransformToOtpEntity()
	workflow.SetOtp(&otp)
	switch payload.RouteType {
	case consts.EmailRouteType.ToString():
		workflow.WithRouteEmail(payload.RouteValue)
	case consts.SMSRouteType.ToString():
		workflow.WithRouteSMS(payload.RouteValue)
	}

	err = workflow.Request(ctx)
	if err != nil {
		data.Message = "ERROR"
	}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)
}
