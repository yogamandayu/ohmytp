package entity

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yogamandayu/ohmytp/storage/repository"
)

// OTPRouteTypeSMS is struct to hold data for otp sms.
type OTPRouteTypeSMS struct {
	ID        string `json:"id"`
	OtpID     string `json:"otp_id"`
	RequestID string `json:"request_id"`
	Phone     string `json:"phone"`
	Timestamp
}

// SetWithOtpRepository is to set entity otp route type sms with repository otp route type sms.
func (o *OTPRouteTypeSMS) SetWithOtpRepository(sms repository.OtpRouteTypeSm) {
	o.ID = sms.ID
	o.OtpID = sms.OtpID
	o.RequestID = sms.RequestID
	o.Phone = sms.Phone.String
	o.CreatedAt = sql.NullTime{
		Time:  sms.CreatedAt.Time,
		Valid: sms.CreatedAt.Valid,
	}
	o.UpdatedAt = sql.NullTime{
		Time:  sms.UpdatedAt.Time,
		Valid: sms.UpdatedAt.Valid,
	}
	o.IsDeleted = sms.IsDeleted.Bool && sms.IsDeleted.Valid
	o.DeletedAt = sql.NullTime{
		Time:  sms.DeletedAt.Time,
		Valid: sms.DeletedAt.Valid,
	}
}

// TransformToOtpRepository is to transform entity otp route type sms to repository otp route type sms.
func (o *OTPRouteTypeSMS) TransformToOtpRepository() (sms repository.OtpRouteTypeSm) {
	sms.ID = o.ID
	sms.RequestID = o.RequestID
	sms.OtpID = o.OtpID
	sms.Phone = pgtype.Text{
		Valid:  true,
		String: o.Phone,
	}
	sms.CreatedAt = pgtype.Timestamptz{
		Time:  o.CreatedAt.Time,
		Valid: o.CreatedAt.Valid,
	}
	sms.UpdatedAt = pgtype.Timestamptz{
		Time:  o.UpdatedAt.Time,
		Valid: o.UpdatedAt.Valid,
	}
	sms.IsDeleted = pgtype.Bool{
		Valid: true,
		Bool:  o.IsDeleted,
	}
	sms.DeletedAt = pgtype.Timestamptz{
		Time:  o.DeletedAt.Time,
		Valid: o.DeletedAt.Valid,
	}
	return
}
