package otp

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yogamandayu/ohmytp/storage/repository"
	"github.com/yogamandayu/ohmytp/tests"
)

type SaveOtpWorkflow struct {
	repository *repository.Queries
	slog       slog.Logger
}

func NewSaveOtpWorkflow(db *pgxpool.Pool, slog *slog.Logger) *SaveOtpWorkflow {
	return &SaveOtpWorkflow{
		repository: repository.New(db),
	}
}

func (w *SaveOtpWorkflow) Save(ctx context.Context) error {
	otp := tests.FakeOtp().TransformToOtpRepository()
	data := repository.SaveOtpParams{
		ID:            otp.ID,
		RequestID:     otp.RequestID,
		RouteType:     otp.RouteType,
		Code:          otp.Code,
		RequestedAt:   otp.RequestedAt,
		ConfirmedAt:   otp.ConfirmedAt,
		ExpiredAt:     otp.ExpiredAt,
		Attempt:       otp.Attempt,
		LastAttemptAt: otp.LastAttemptAt,
		ResendAttempt: otp.ResendAttempt,
		ResendAt:      otp.ResendAt,
		IpAddress:     otp.IpAddress,
		UserAgent:     otp.UserAgent,
	}
	_, err := w.repository.SaveOtp(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
