package otp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/response"
	"github.com/yogamandayu/ohmytp/internal/requester"
	"github.com/yogamandayu/ohmytp/internal/workflow/otp"
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
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusBadRequest).AsJSON(w)
		return
	}
	rq := requester.NewRequester().SetMetadataFromREST(r)

	th := throttle.NewThrottle(h.app.RedisAPI, "request_otp", payload.Identifier).SetThresholds([]throttle.Threshold{
		{
			MaxAttempt:      1,
			WaitingDuration: 2 * time.Minute,
		}, {
			MaxAttempt:      1,
			WaitingDuration: 5 * time.Minute,
		}, {
			MaxAttempt:      1,
			WaitingDuration: 15 * time.Minute,
		}, {
			MaxAttempt:      1,
			WaitingDuration: 24 * time.Hour,
		},
	})
	ok, _ := th.IsAllowed(ctx)
	if !ok {
		err = fmt.Errorf("otp.error.request_otp.throttled_until:%s", th.WaitUntil().Format(time.RFC3339))
		h.app.Log.Error(err.Error())
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusTooManyRequests).AsJSON(w)
		return
	}

	workflow := otp.NewRequestOtpWorkflow(rq, h.app)
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
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusInternalServerError).AsJSON(w)
		return
	}
	response.NewHTTPSuccessResponse(RequestOtpResponseContract{
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
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusBadRequest).AsJSON(w)
		return
	}

	rq := requester.NewRequester().SetMetadataFromREST(r)

	th := throttle.NewThrottle(h.app.RedisAPI, "confirm_otp", payload.Identifier).SetThresholds([]throttle.Threshold{
		{
			MaxAttempt:      3,
			WaitingDuration: 2 * time.Minute,
		}, {
			MaxAttempt:      3,
			WaitingDuration: 5 * time.Minute,
		}, {
			MaxAttempt:      3,
			WaitingDuration: 15 * time.Minute,
		}, {
			MaxAttempt:      3,
			WaitingDuration: 24 * time.Hour,
		},
	})

	ok, _ := th.IsAllowed(ctx)
	if !ok {
		err = fmt.Errorf("otp.error.confirm_otp.throttled_until:%s", th.WaitUntil().Format(time.RFC3339))
		h.app.Log.Error(err.Error())
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusTooManyRequests).AsJSON(w)
		return
	}

	workflow := otp.NewConfirmOtpWorkflow(rq, h.app)
	otpEntity := payload.TransformToOtpEntity()
	workflow.SetOtp(&otpEntity)
	err = workflow.Confirm(ctx)
	if err != nil {
		h.app.Log.Error(err.Error())
		response.NewHTTPFailedResponse("ERR101", err, "Error").WithStatusCode(http.StatusUnprocessableEntity).AsJSON(w)
		return
	}
	response.NewHTTPSuccessResponse(nil, "Success").WithStatusCode(http.StatusCreated).AsJSON(w)
}
