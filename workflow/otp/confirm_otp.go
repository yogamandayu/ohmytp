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

// SetOtp is to set entity.Otp to RequestOtpWorkflow.
func (c *ConfirmOtpWorkflow) SetOtp(otp *entity.Otp) *ConfirmOtpWorkflow {
	c.Otp = otp
	return c
}

func (c *ConfirmOtpWorkflow) Confirm() error {

	return nil
}
