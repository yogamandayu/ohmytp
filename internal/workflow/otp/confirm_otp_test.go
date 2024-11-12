package otp_test

import (
	"context"
	"testing"
	"time"

	"github.com/yogamandayu/ohmytp/internal/domain/entity"
	"github.com/yogamandayu/ohmytp/internal/requester"
	tests2 "github.com/yogamandayu/ohmytp/internal/tests"
	otp2 "github.com/yogamandayu/ohmytp/internal/workflow/otp"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/consts"
	"github.com/yogamandayu/ohmytp/util"
)

func TestConfirmOtp(t *testing.T) {
	ts := tests2.NewTestSuite()
	ts.LoadApp()
	defer t.Cleanup(ts.Clean)

	scenarios := []struct {
		id           string
		description  string
		otp          entity.Otp
		totalAttempt int
		expiration   time.Duration
		isErr        bool
		errMsg       string
	}{
		{
			id:          uuid.NewString(),
			description: "Positive case to confirm otp with valid otp",
			otp: entity.Otp{
				Identifier: uuid.NewString(),
				Purpose:    "TEST",
			},
			totalAttempt: 1,
			isErr:        false,
		},
		{
			id:          uuid.NewString(),
			description: "Negative case to confirm otp with invalid otp",
			otp: entity.Otp{
				Identifier: uuid.NewString(),
				Purpose:    "TEST",
				Code:       "INVALID",
			},
			totalAttempt: 1,
			isErr:        true,
			errMsg:       "otp.error.confirm_otp.invalid_otp_code",
		},
		{
			id:          uuid.NewString(),
			description: "Negative case to confirm otp with invalid otp and try again after reach max attempt",
			otp: entity.Otp{
				Identifier: uuid.NewString(),
				Purpose:    "TEST",
				Code:       "INVALID",
			},
			totalAttempt: util.GetEnvAsInt("MAX_CONFIRM_OTP_ATTEMPT", 3) + 1,
			isErr:        true,
			errMsg:       "otp.error.confirm_otp.max_attempt_reached",
		},
		{
			id:          uuid.NewString(),
			description: "Negative case to confirm otp with valid otp but already expired",
			otp: entity.Otp{
				Identifier: uuid.NewString(),
				Purpose:    "TEST",
			},
			expiration:   1 * time.Second,
			totalAttempt: 1,
			isErr:        true,
			errMsg:       "otp.error.confirm_otp.otp_is_expired",
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			rqs := requester.NewRequester().SetMetadataFromREST(tests2.FakeHTTPRequest())
			rqs.Metadata.RequestID = uuid.NewString()
			requestWorkflow := otp2.NewRequestOtpWorkflow(rqs, ts.App)
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
				confirmWorkflow := otp2.NewConfirmOtpWorkflow(rqs, ts.App)
				if scenario.otp.Code == "" {
					scenario.otp.Code = resOtpEntity.Code
				}

				confirmWorkflow.SetOtp(&scenario.otp)

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
