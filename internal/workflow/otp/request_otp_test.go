package otp_test

import (
	"context"
	"testing"

	"github.com/yogamandayu/ohmytp/internal/domain/entity"
	"github.com/yogamandayu/ohmytp/internal/requester"
	"github.com/yogamandayu/ohmytp/internal/tests"
	"github.com/yogamandayu/ohmytp/internal/workflow/otp"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/consts"
)

func TestRequestOtp(t *testing.T) {
	ts := tests.NewTestSuite()
	ts.LoadApp()
	defer t.Cleanup(ts.Clean)

	scenarios := []struct {
		id          string
		description string
		routeType   string
		email       string
		phone       string
		isErr       bool
	}{
		{
			id:          uuid.NewString(),
			description: "Positive case to request otp email with valid data",
			routeType:   consts.EmailRouteType.ToString(),
			email:       "example@example.com",
		},
		{
			id:          uuid.NewString(),
			description: "Positive case to request otp sms with valid data",
			routeType:   consts.SMSRouteType.ToString(),
			phone:       "1234567890",
		},
		{
			id:          uuid.NewString(),
			description: "Negative case to request otp sms with invalid route type",
			routeType:   "invalid-route",
			phone:       "1234567890",
			isErr:       true,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			rqs := requester.NewRequester().SetMetadataFromREST(tests.FakeHTTPRequest())
			workflow := otp.NewRequestOtpWorkflow(rqs, ts.App)
			workflow.SetOtp(&entity.Otp{
				ID:        uuid.NewString(),
				RequestID: uuid.NewString(),
				RouteType: scenario.routeType,
				IPAddress: "127.0.0.1",
				UserAgent: "Test-Func",
			})
			switch scenario.routeType {
			case consts.EmailRouteType.ToString():
				_ = workflow.WithRouteEmail(scenario.email)
			case consts.SMSRouteType.ToString():
				_ = workflow.WithRouteSMS(scenario.phone)
			}

			_, err := workflow.Request(context.Background())
			if !scenario.isErr {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
