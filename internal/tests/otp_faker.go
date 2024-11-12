package tests

import (
	"database/sql"
	entity2 "github.com/yogamandayu/ohmytp/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

// FakeOtp return faker otp entity.
func FakeOtp() *entity2.Otp {
	return &entity2.Otp{
		ID:         uuid.NewString(),
		RequestID:  uuid.NewString(),
		Identifier: uuid.NewString(),
		RouteType:  "EMAIL",
		Code:       "12345",
		Purpose:    "TEST",
		RequestedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ConfirmedAt: sql.NullTime{},
		ExpiredAt: sql.NullTime{
			Time:  time.Now().Add(3 * time.Minute),
			Valid: true,
		},
		Attempt:       0,
		LastAttemptAt: sql.NullTime{},
		ResendAttempt: 0,
		ResendAt:      sql.NullTime{},
		IPAddress:     "127.0.0.1",
		UserAgent:     "",
		Timestamp: entity2.Timestamp{
			CreatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		},
	}
}
