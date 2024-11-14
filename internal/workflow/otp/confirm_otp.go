package otp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/domain/entity"
	"github.com/yogamandayu/ohmytp/internal/requester"
	"github.com/yogamandayu/ohmytp/internal/storage/cache"
	"github.com/yogamandayu/ohmytp/internal/storage/repository"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yogamandayu/ohmytp/util"
)

// ConfirmOtpWorkflow is request OTP workflow.
type ConfirmOtpWorkflow struct {
	Otp *entity.Otp

	App       *app.App
	Requester *requester.Requester
}

// SetOtp is to set entity.Otp to ConfirmOtpWorkflow.
func (c *ConfirmOtpWorkflow) SetOtp(otp *entity.Otp) *ConfirmOtpWorkflow {
	c.Otp = otp
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

	if c.App.RedisAPI != nil {
		otpCache := cache.NewOTPCache(c.App.RedisAPI)
		otpEntity = otpCache.GetOTP(ctx, c.Requester.Metadata.RequestID)
	}
	if otpEntity.Code == "" {
		var findOtpRes repository.FindOtpByIdentifierAndPurposeRow
		findOtpParams := repository.FindOtpByIdentifierAndPurposeParams{
			Identifier: pgtype.Text{
				String: c.Otp.Identifier,
				Valid:  true,
			},
			Purpose: pgtype.Text{
				String: c.Otp.Purpose,
				Valid:  true,
			},
		}
		findOtpRes, err = c.App.DBRepository.FindOtpByIdentifierAndPurpose(ctx, findOtpParams)
		if err != nil {
			return err
		}
		otpEntity.SetWithFindOtpRepositoryByIdentifierAndPurpose(findOtpRes)
	}

	if otpEntity.ConfirmedAt.Valid {
		err = errors.New("otp.error.confirm_otp.otp_already_confirmed")
		c.updateAttempt(ctx, otpEntity, err)
		return err
	}

	if int(otpEntity.Attempt) >= util.GetEnvAsInt("MAX_CONFIRM_OTP_ATTEMPT", 3) {
		err = errors.New("otp.error.confirm_otp.max_attempt_reached")
		c.updateAttempt(ctx, otpEntity, err)
		return err
	}

	if time.Now().After(otpEntity.ExpiredAt.Time) {
		err = errors.New("otp.error.confirm_otp.otp_is_expired")
		c.updateAttempt(ctx, otpEntity, err)
		return err
	}

	if otpEntity.Code != c.Otp.Code {
		err = errors.New("otp.error.confirm_otp.invalid_otp_code")
		c.updateAttempt(ctx, otpEntity, err)
		return err
	}
	c.updateAttempt(ctx, otpEntity, nil)
	return nil
}

// updateAttempt is a post process after confirming otp to update otp.
func (c *ConfirmOtpWorkflow) updateAttempt(ctx context.Context, otpEntity entity.Otp, err error) {
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

	if c.App.RedisAPI != nil {
		otpCache := cache.NewOTPCache(c.App.RedisAPI)
		otpCache.InvalidateOTP(ctx, c.Requester.Metadata.RequestID)
		if err != nil {
			var otp entity.Otp
			otp.SetWithUpdateOtpAttemptRepository(updateOtpAttemptRes)
			otpCache.SetOTP(ctx, c.Requester.Metadata.RequestID, otp, time.Duration(updateOtpAttemptRes.ExpiredAt.Time.Second()-time.Now().Second())+(30*time.Second))
		}
	}
}
