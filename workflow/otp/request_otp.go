package otp

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/consts"
	"github.com/yogamandayu/ohmytp/domain/entity"
	"github.com/yogamandayu/ohmytp/requester"
	"github.com/yogamandayu/ohmytp/storage/repository"
)

// RequestOtpWorkflow is request OTP workflow.
type RequestOtpWorkflow struct {
	Otp            *entity.Otp
	RouteTypeEmail *entity.OTPRouteTypeEmail
	RouteTypeSMS   *entity.OTPRouteTypeSMS

	App       *app.App
	Requester *requester.Requester
}

func (r *RequestOtpWorkflow) SetOtp(otp *entity.Otp) *RequestOtpWorkflow {
	r.Otp = otp
	return r
}

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
func NewRequestOtpWorkflow(requester *requester.Requester, app *app.App) *RequestOtpWorkflow {
	return &RequestOtpWorkflow{
		App:       app,
		Requester: requester,
	}
}

// Request is requesting OTP.
func (r *RequestOtpWorkflow) Request(ctx context.Context) error {

	generatedOtp, _ := rand.Int(rand.Reader, big.NewInt(99999))
	uid, _ := uuid.NewV7()

	tx, err := r.App.DB.Begin(ctx)
	if err != nil {
		return err
	}

	dataOtp := repository.SaveOtpParams{
		ID:        uid.String(),
		RequestID: r.Requester.Metadata.RequestID,
		RouteType: pgtype.Text{
			Valid:  true,
			String: r.Otp.RouteType,
		},
		Code: pgtype.Text{
			Valid:  true,
			String: strconv.Itoa(int(generatedOtp.Int64())),
		},
		RequestedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiredAt: pgtype.Timestamptz{
			Time:  time.Now().Add(2 * time.Minute),
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
	_, err = r.App.DBRepository.WithTx(tx).SaveOtp(ctx, dataOtp)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	switch r.Otp.RouteType {
	case "SMS":
		dataSMS := repository.SaveOtpRouteTypeSMSParams{
			ID:        uid.String(),
			RequestID: dataOtp.RequestID,
			OtpID:     dataOtp.ID,
			Phone: pgtype.Text{
				String: r.RouteTypeSMS.Phone,
				Valid:  true,
			},
		}
		_, err = r.App.DBRepository.WithTx(tx).SaveOtpRouteTypeSMS(ctx, dataSMS)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	case "EMAIL":
		dataEmail := repository.SaveOtpRouteTypeEmailParams{
			ID:        uid.String(),
			RequestID: dataOtp.RequestID,
			OtpID:     dataOtp.ID,
			Email: pgtype.Text{
				String: r.RouteTypeEmail.Email,
				Valid:  true,
			},
		}
		_, err = r.App.DBRepository.WithTx(tx).SaveOtpRouteTypeEmail(ctx, dataEmail)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	default:
		tx.Rollback(ctx)
		return errors.New("invalid route type")
	}

	tx.Commit(ctx)
	return nil
}
