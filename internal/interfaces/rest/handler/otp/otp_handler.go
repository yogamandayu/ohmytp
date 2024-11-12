package otp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	response2 "github.com/yogamandayu/ohmytp/internal/interfaces/rest/response"
	"github.com/yogamandayu/ohmytp/internal/requester"
	otp2 "github.com/yogamandayu/ohmytp/internal/workflow/otp"
	"github.com/yogamandayu/ohmytp/pkg/throttle"

	"encoding/json"

	"github.com/yogamandayu/ohmytp/consts"
)

// Request is request otp request handler.
func (h *Handler) Request(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)

	var payload RequestOtpRequestContract
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		h.app.Log.Warn(err.Error())
		response2.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusBadRequest).AsJSON(w)
		return
	}
	rq := requester.NewRequester().SetMetadataFromREST(r)

	th := throttle.NewThrottle(h.app.Redis, "request_otp", rq.Metadata.RequestID).SetThresholds([]throttle.Threshold{
		{
			MaxAttempt:      3,
			WaitingDuration: 30 * time.Second,
		}, {
			MaxAttempt:      5,
			WaitingDuration: 60 * time.Second,
		},
	})
	ok, _ := th.IsAllowed(ctx)
	if !ok {
		err = errors.New(fmt.Sprintf("otp.error.request_otp.throttled:%s", th.WaitUntil().Format(time.RFC3339)))
		h.app.Log.Error(err.Error())
		response2.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusTooManyRequests).AsJSON(w)
		return
	}

	workflow := otp2.NewRequestOtpWorkflow(rq, h.app)
	otpEntity := payload.TransformToOtpEntity()

	workflow.SetOtp(&otpEntity).SetOtpExpiration(time.Duration(payload.Expiration) * time.Second).SetOtpLength(payload.Length)
	switch otpEntity.RouteType {
	case consts.EmailRouteType.ToString():
		_ = workflow.WithRouteEmail(payload.RouteValue)
	case consts.SMSRouteType.ToString():
		_ = workflow.WithRouteSMS(payload.RouteValue)
	}
	resOtpEntity, err := workflow.Request(ctx)
	if err != nil {
		h.app.Log.Error(err.Error())
		response2.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusInternalServerError).AsJSON(w)
		return
	}
	response2.NewHTTPSuccessResponse(RequestOtpResponseContract{
		ExpiredAt: resOtpEntity.ExpiredAt.Time.Format(time.RFC3339),
	}, "Success").WithStatusCode(http.StatusCreated).AsJSON(w)
}

// Confirm is request otp confirm handler.
func (h *Handler) Confirm(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)

	var payload ConfirmOtpRequestContract
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		h.app.Log.Warn(err.Error())
		response2.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusBadRequest).AsJSON(w)
		return
	}

	rq := requester.NewRequester().SetMetadataFromREST(r)

	th := throttle.NewThrottle(h.app.Redis, "confirm_otp", rq.Metadata.RequestID).SetThresholds([]throttle.Threshold{
		{
			MaxAttempt:      3,
			WaitingDuration: 30 * time.Second,
		}, {
			MaxAttempt:      5,
			WaitingDuration: 60 * time.Second,
		},
	})

	ok, _ := th.IsAllowed(ctx)
	if !ok {
		err = errors.New(fmt.Sprintf("otp.error.confirm_otp.throttled:%s", th.WaitUntil().Format(time.RFC3339)))
		h.app.Log.Error(err.Error())
		response2.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusTooManyRequests).AsJSON(w)
		return
	}

	workflow := otp2.NewConfirmOtpWorkflow(rq, h.app)
	otpEntity := payload.TransformToOtpEntity()
	workflow.SetOtp(&otpEntity)
	err = workflow.Confirm(ctx)
	if err != nil {
		h.app.Log.Error(err.Error())
		response2.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusUnprocessableEntity).AsJSON(w)
		return
	}
	response2.NewHTTPSuccessResponse(nil, "Success").WithStatusCode(http.StatusCreated).AsJSON(w)
}
