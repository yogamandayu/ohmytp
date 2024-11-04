package otp

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/domain/entity"
	"github.com/yogamandayu/ohmytp/requester"
	"github.com/yogamandayu/ohmytp/storage/cache"
	"github.com/yogamandayu/ohmytp/storage/repository"
	"github.com/yogamandayu/ohmytp/util"
	"time"
)

// ConfirmOtpWorkflow is request OTP workflow.
type ConfirmOtpWorkflow struct {
	OtpCode string

	App       *app.App
	Requester *requester.Requester
}

// SetOtpCode is to set otp code to RequestOtpWorkflow.
func (c *ConfirmOtpWorkflow) SetOtpCode(code string) *ConfirmOtpWorkflow {
	c.OtpCode = code
	return c
}

// NewConfirmOtpWorkflow is a constructor.
func NewConfirmOtpWorkflow(rqs *requester.Requester, app *app.App) *ConfirmOtpWorkflow {
	return &ConfirmOtpWorkflow{
		App:       app,
		Requester: rqs,
	}
}

// Confirm is to confirm otp. if otp is failed to confirm, attempt is increased.
func (c *ConfirmOtpWorkflow) Confirm(ctx context.Context) error {
	var err error
	var otpEntity entity.Otp

	defer func(err error) {
		c.UpdateOTP(ctx, otpEntity, err)
	}(err)

	if c.App.Redis != nil {
		otpCache := cache.NewOTPCache(c.App.Redis)
		otpEntity = otpCache.GetOTP(ctx, c.Requester.Metadata.RequestID)
	}
	if otpEntity.Code == "" {
		var findOtpRes repository.FindOtpByRequestIDRow
		findOtpRes, err = c.App.DBRepository.FindOtpByRequestID(ctx, c.Requester.Metadata.RequestID)
		if err != nil {
			return err
		}
		otpEntity.SetWithFindOtpRepositoryByRequestID(findOtpRes)
	}

	if int(otpEntity.Attempt) >= util.GetEnvAsInt("MAX_CONFIRM_OTP_ATTEMPT", 3) {
		return errors.New("otp.error.confirm_otp.max_attempt_reached")
	}

	if time.Now().After(otpEntity.ExpiredAt.Time) {
		return errors.New("otp.error.confirm_otp.otp_is_expired")
	}

	if otpEntity.Code != c.OtpCode {
		return errors.New("otp.error.confirm_otp.invalid_otp_code")
	}

	return nil
}

// UpdateOTP is a post process after confirming otp to update otp.
func (c *ConfirmOtpWorkflow) UpdateOTP(ctx context.Context, otpEntity entity.Otp, err error) {
	updateOtpParams := repository.UpdateOtpAttemptParams{
		ID: otpEntity.ID,
		Attempt: pgtype.Int2{
			Int16: int16(otpEntity.Attempt + 1),
			Valid: true,
		},
		LastAttemptAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}
	if err == nil {
		updateOtpParams.ConfirmedAt = pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		}
	}

	updateOtpAttemptRes, errUpdate := c.App.DBRepository.UpdateOtpAttempt(ctx, updateOtpParams)
	if errUpdate != nil {
		c.App.Log.Error(fmt.Sprintf("error update otp attempt when confirming, err: %v", errUpdate))
	}

	if c.App.Redis != nil {
		otpCache := cache.NewOTPCache(c.App.Redis)
		otpCache.InvalidateOTP(ctx, c.Requester.Metadata.RequestID)
		if err != nil {
			var otp entity.Otp
			otp.SetWithUpdateOtpAttemptRepository(updateOtpAttemptRes)
			otpCache.SetOTP(ctx, c.Requester.Metadata.RequestID, otp, time.Duration(updateOtpAttemptRes.ExpiredAt.Time.Second()-time.Now().Second())+(30*time.Second))
		}
	}
	return
}
