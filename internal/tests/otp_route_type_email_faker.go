package tests

import (
	"database/sql"
	"time"

	entity2 "github.com/yogamandayu/ohmytp/internal/domain/entity"

	"github.com/google/uuid"
)

// FakeOtpRouteTypeEmail return faker otp route type email entity.
func FakeOtpRouteTypeEmail() *entity2.OTPRouteTypeEmail {
	return &entity2.OTPRouteTypeEmail{
		ID:        uuid.NewString(),
		OtpID:     uuid.NewString(),
		RequestID: uuid.NewString(),
		Email:     "example@example.com",
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
