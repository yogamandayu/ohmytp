package otp

import (
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/domain/entity"
	"github.com/yogamandayu/ohmytp/requester"
)

// ConfirmOtpWorkflow is request OTP workflow.
type ConfirmOtpWorkflow struct {
	Otp *entity.Otp

	App       *app.App
	Requester *requester.Requester
}
