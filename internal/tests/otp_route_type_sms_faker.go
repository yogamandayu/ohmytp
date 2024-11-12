package tests

import (
	"database/sql"
	entity2 "github.com/yogamandayu/ohmytp/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

// FakeOtpRouteTypeSMS return faker otp route type sms entity.
func FakeOtpRouteTypeSMS() *entity2.OTPRouteTypeSMS {
	return &entity2.OTPRouteTypeSMS{
		ID:        uuid.NewString(),
		OtpID:     uuid.NewString(),
		RequestID: uuid.NewString(),
		Phone:     "0987654321",
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
