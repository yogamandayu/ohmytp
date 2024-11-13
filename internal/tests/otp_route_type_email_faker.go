package tests

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/yogamandayu/ohmytp/internal/domain/entity"
)

// FakeOtpRouteTypeEmail return faker otp route type email entity.
func FakeOtpRouteTypeEmail() *entity.OTPRouteTypeEmail {
	return &entity.OTPRouteTypeEmail{
		ID:        uuid.NewString(),
		OtpID:     uuid.NewString(),
		RequestID: uuid.NewString(),
		Email:     "example@example.com",
		Timestamp: entity.Timestamp{
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
