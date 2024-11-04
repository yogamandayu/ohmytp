package otp_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/consts"
	"github.com/yogamandayu/ohmytp/domain/entity"
	"github.com/yogamandayu/ohmytp/requester"
	"github.com/yogamandayu/ohmytp/tests"
	"github.com/yogamandayu/ohmytp/util"
	"github.com/yogamandayu/ohmytp/workflow/otp"
	"testing"
	"time"
)

func TestConfirmOtp(t *testing.T) {
	ts := tests.NewTestSuite()
	ts.LoadApp()
	defer t.Cleanup(ts.Clean)

	scenarios := []struct {
		id           string
		description  string
		requestID    string
		totalAttempt int
		expiration   time.Duration
		isErr        bool
		errMsg       string
		code         string
	}{
		{
			id:           uuid.NewString(),
			description:  "Positive case to confirm otp with valid otp",
			requestID:    uuid.NewString(),
			totalAttempt: 1,
			isErr:        false,
		},
		{
			id:           uuid.NewString(),
			description:  "Negative case to confirm otp with invalid otp",
			requestID:    uuid.NewString(),
			totalAttempt: 1,
			isErr:        true,
			errMsg:       "otp.error.confirm_otp.invalid_otp_code",
			code:         "INVALID",
		},
		{
			id:           uuid.NewString(),
			description:  "Negative case to confirm otp with invalid otp and try again after reach max attempt",
			requestID:    uuid.NewString(),
			totalAttempt: util.GetEnvAsInt("MAX_CONFIRM_OTP_ATTEMPT", 3) + 1,
			isErr:        true,
			errMsg:       "otp.error.confirm_otp.max_attempt_reached",
			code:         "INVALID",
		},
		{
			id:           uuid.NewString(),
			description:  "Negative case to confirm otp with valid otp but already expired",
			requestID:    uuid.NewString(),
			expiration:   1 * time.Second,
			totalAttempt: 1,
			isErr:        true,
			errMsg:       "otp.error.confirm_otp.otp_is_expired",
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			rqs := requester.NewRequester().SetMetadataFromREST(tests.FakeHTTPRequest())
			rqs.Metadata.RequestID = scenario.requestID
			requestWorkflow := otp.NewRequestOtpWorkflow(rqs, ts.App)
			requestWorkflow.SetOtp(&entity.Otp{
				RouteType: consts.EmailRouteType.ToString(),
				Purpose:   "TEST",
			})
			if scenario.expiration != 0 {
				requestWorkflow.SetOtpExpiration(scenario.expiration)
			}
			_ = requestWorkflow.WithRouteEmail("example@example.com")

			resOtpEntity, err := requestWorkflow.Request(context.Background())
			require.NoError(t, err)

			for i := range scenario.totalAttempt {
				confirmWorkflow := otp.NewConfirmOtpWorkflow(rqs, ts.App)
				confirmWorkflow.SetOtpCode(resOtpEntity.Code)
				if scenario.code != "" {
					confirmWorkflow.SetOtpCode(scenario.code)
				}

				if scenario.expiration != 0 {
					duration := scenario.expiration + (1 * time.Second)
					time.Sleep(duration)
				}

				err = confirmWorkflow.Confirm(context.Background())
				if !scenario.isErr {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					if i == scenario.totalAttempt-1 {
						assert.Equal(t, scenario.errMsg, err.Error())
					}
					if scenario.expiration != 0 {
						assert.Equal(t, scenario.errMsg, err.Error())
					}
				}
			}
		})
	}
}
