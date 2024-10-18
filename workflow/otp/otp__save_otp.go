package otp

import (
	"context"
	"crypto/rand"
	"log/slog"
	"math/big"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yogamandayu/ohmytp/domain/entity"
	"github.com/yogamandayu/ohmytp/storage/repository"
)

// RequestOtpWorkflow is request OTP workflow.
type RequestOtpWorkflow struct {
	repository *repository.Queries
	slog       *slog.Logger
}

// NewRequestOtpWorkflow is a constructor.
func NewRequestOtpWorkflow(db *pgxpool.Pool, slog *slog.Logger) *RequestOtpWorkflow {
	return &RequestOtpWorkflow{
		repository: repository.New(db),
		slog:       slog,
	}
}

// Request is requesting OTP.
func (w *RequestOtpWorkflow) Request(ctx context.Context, otp entity.Otp) error {
	otpRepo := otp.TransformToOtpRepository()

	generatedOtp, _ := rand.Int(rand.Reader, big.NewInt(99999))
	data := repository.SaveOtpParams{
		ID:        uuid.NewString(),
		RequestID: ctx.Value("X-Request-ID").(string),
		RouteType: otpRepo.RouteType,
		Code: pgtype.Text{
			Valid:  true,
			String: strconv.Itoa(int(generatedOtp.Int64())),
		},
		RequestedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		ConfirmedAt: otpRepo.ConfirmedAt,
		ExpiredAt: pgtype.Timestamptz{
			Time:  time.Now().Add(2 * time.Minute),
			Valid: true,
		},
		IpAddress: otpRepo.IpAddress,
		UserAgent: otpRepo.UserAgent,
	}
	_, err := w.repository.SaveOtp(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
