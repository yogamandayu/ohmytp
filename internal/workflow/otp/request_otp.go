package otp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yogamandayu/ohmytp/pkg/worker"

	"github.com/yogamandayu/ohmytp/pkg/telegram"

	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/domain/entity"
	"github.com/yogamandayu/ohmytp/internal/requester"
	"github.com/yogamandayu/ohmytp/internal/storage/cache"
	"github.com/yogamandayu/ohmytp/internal/storage/repository"

	"github.com/yogamandayu/ohmytp/util"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yogamandayu/ohmytp/consts"
)

// RequestOtpWorkflow is request OTP workflow.
type RequestOtpWorkflow struct {
	Otp            *entity.Otp
	RouteTypeEmail *entity.OTPRouteTypeEmail
	RouteTypeSMS   *entity.OTPRouteTypeSMS
	Expiration     time.Duration
	OtpLength      uint16

	App       *app.App
	Requester *requester.Requester
}

// SetOtp is to set entity.Otp to RequestOtpWorkflow.
func (r *RequestOtpWorkflow) SetOtp(otp *entity.Otp) *RequestOtpWorkflow {
	r.Otp = otp
	return r
}

// SetOtpLength is to set generated OTP length.
func (r *RequestOtpWorkflow) SetOtpLength(length uint16) *RequestOtpWorkflow {
	r.OtpLength = length
	return r
}

// SetOtpExpiration is to set OTP expiration time.
func (r *RequestOtpWorkflow) SetOtpExpiration(expiration time.Duration) *RequestOtpWorkflow {
	r.Expiration = expiration
	return r
}

// WithRouteEmail is to set request OTP route type to email.
func (r *RequestOtpWorkflow) WithRouteEmail(email string) error {
	if r.Otp == nil {
		return errors.New("missing otp")
	}
	r.Otp.RouteType = consts.EmailRouteType.ToString()
	uid, _ := uuid.NewV7()
	r.RouteTypeEmail = &entity.OTPRouteTypeEmail{
		ID:        uid.String(),
		OtpID:     r.Otp.ID,
		RequestID: r.Otp.RequestID,
		Email:     email,
	}
	return nil
}

// WithRouteSMS is to set request OTP route type to SMS.
func (r *RequestOtpWorkflow) WithRouteSMS(phone string) error {
	if r.Otp == nil {
		return errors.New("missing otp")
	}
	r.Otp.RouteType = consts.SMSRouteType.ToString()
	uid, _ := uuid.NewV7()
	r.RouteTypeSMS = &entity.OTPRouteTypeSMS{
		ID:        uid.String(),
		OtpID:     r.Otp.ID,
		RequestID: r.Otp.RequestID,
		Phone:     phone,
	}
	return nil
}

// NewRequestOtpWorkflow is a constructor.
func NewRequestOtpWorkflow(rqs *requester.Requester, app *app.App) *RequestOtpWorkflow {
	return &RequestOtpWorkflow{
		App:        app,
		Requester:  rqs,
		OtpLength:  5,
		Expiration: 2 * time.Minute,
	}
}

// Request is requesting OTP.
func (r *RequestOtpWorkflow) Request(ctx context.Context) (otpEntity entity.Otp, err error) {

	generatedOtp := util.RandomStringWithSample(int(r.OtpLength), "0123456789")
	uid, _ := uuid.NewV7()

	tx, err := r.App.DB.Begin(ctx)
	if err != nil {
		return entity.Otp{}, err
	}

	saveOtpParams := repository.SaveOtpParams{
		ID:        uid.String(),
		RequestID: r.Requester.Metadata.RequestID,
		Identifier: pgtype.Text{
			String: r.Otp.Identifier,
			Valid:  true,
		},
		RouteType: pgtype.Text{
			Valid:  true,
			String: r.Otp.RouteType,
		},
		Purpose: pgtype.Text{
			Valid:  true,
			String: r.Otp.Purpose,
		},
		Code: pgtype.Text{
			Valid:  true,
			String: generatedOtp,
		},
		RequestedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiredAt: pgtype.Timestamptz{
			Time:  time.Now().Add(r.Expiration),
			Valid: true,
		},
		IpAddress: pgtype.Text{
			Valid:  true,
			String: r.Requester.Metadata.IPAddress,
		},
		UserAgent: pgtype.Text{
			Valid:  true,
			String: r.Requester.Metadata.UserAgent,
		},
	}
	saveOtpRes, err := r.App.DBRepository.WithTx(tx).SaveOtp(ctx, saveOtpParams)
	if err != nil {
		_ = tx.Rollback(ctx)
		return entity.Otp{}, err
	}

	switch r.Otp.RouteType {
	case "SMS":
		dataSMS := repository.SaveOtpRouteTypeSMSParams{
			ID:        uid.String(),
			RequestID: r.Requester.Metadata.RequestID,
			OtpID:     saveOtpRes.ID,
			Phone: pgtype.Text{
				String: r.RouteTypeSMS.Phone,
				Valid:  true,
			},
		}
		_, err = r.App.DBRepository.WithTx(tx).SaveOtpRouteTypeSMS(ctx, dataSMS)
		if err != nil {
			_ = tx.Rollback(ctx)
			return entity.Otp{}, err
		}
	case "EMAIL":
		dataEmail := repository.SaveOtpRouteTypeEmailParams{
			ID:        uid.String(),
			RequestID: r.Requester.Metadata.RequestID,
			OtpID:     saveOtpRes.ID,
			Email: pgtype.Text{
				String: r.RouteTypeEmail.Email,
				Valid:  true,
			},
		}
		_, err = r.App.DBRepository.WithTx(tx).SaveOtpRouteTypeEmail(ctx, dataEmail)
		if err != nil {
			_ = tx.Rollback(ctx)
			return entity.Otp{}, err
		}
	default:
		_ = tx.Rollback(ctx)
		return entity.Otp{}, errors.New("otp.error.request_otp.invalid_route_type")
	}
	_ = tx.Commit(ctx)
	otpEntity.SetWithOtpRepository(saveOtpRes)
	if r.App.RedisAPI != nil {
		otpCache := cache.NewOTPCache(r.App.RedisAPI)
		otpCache.SetOTP(ctx, r.Requester.Metadata.RequestID, otpEntity, r.Expiration+(30*time.Second))
	}
	if r.App.Config.TelegramBot != nil {
		err = r.SendOTPToTelegram(otpEntity)
		if err != nil {
			r.App.Log.Error("error sending otp to telegram", "error", err)
		}
	}
	return
}

func (r *RequestOtpWorkflow) SendOTPToTelegram(otpEntity entity.Otp) error {

	workerPool := worker.NewWorker(r.App.RedisWorkerNotification)
	workerPool.AsProducer(&worker.ProducerConfig{
		MaxRetry: 3,
	})
	workerPool.Produce("worker:notification", []byte("Test Worker!"))

	var message string
	switch r.Otp.RouteType {
	case "SMS":
		message = fmt.Sprintf("Your OhMyTP via sms is %s", otpEntity.Code)
	case "EMAIL":
		message = fmt.Sprintf("Your OhMyTP via email is %s", otpEntity.Code)
	}
	bot := telegram.NewTelegramBot(r.App.Log, r.App.Config.TelegramBot.Config)
	return bot.SendMessage(message)
}
