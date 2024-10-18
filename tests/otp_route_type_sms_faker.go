package tests

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/yogamandayu/ohmytp/domain/entity"
)

// FakeOtpRouteTypeSMS return faker otp route type sms entity.
func FakeOtpRouteTypeSMS() *entity.OTPRouteTypeSMS {
	return &entity.OTPRouteTypeSMS{
		ID:        uuid.NewString(),
		OtpID:     uuid.NewString(),
		RequestID: uuid.NewString(),
		Phone:     "0987654321",
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
