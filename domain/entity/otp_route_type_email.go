package entity

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yogamandayu/ohmytp/storage/repository"
)

// OTPRouteTypeEmail is struct to hold data for otp email.
type OTPRouteTypeEmail struct {
	ID        string `json:"id"`
	OtpID     string `json:"otp_id"`
	RequestID string `json:"request_id"`
	Email     string `json:"email"`
	Timestamp
}

// SetWithOtpRepository is to set entity otp route type email with repository otp route type email.
func (o *OTPRouteTypeEmail) SetWithOtpRepository(email repository.OtpRouteTypeEmail) {
	o.ID = email.ID
	o.OtpID = email.OtpID
	o.RequestID = email.RequestID
	o.Email = email.Email.String
	o.CreatedAt = sql.NullTime{
		Time:  email.CreatedAt.Time,
		Valid: email.CreatedAt.Valid,
	}
	o.UpdatedAt = sql.NullTime{
		Time:  email.UpdatedAt.Time,
		Valid: email.UpdatedAt.Valid,
	}
	o.IsDeleted = email.IsDeleted.Bool && email.IsDeleted.Valid
	o.DeletedAt = sql.NullTime{
		Time:  email.DeletedAt.Time,
		Valid: email.DeletedAt.Valid,
	}
}

// TransformToOtpRepository is to transform entity otp route type email to repository otp route type email.
func (o *OTPRouteTypeEmail) TransformToOtpRepository() (email repository.OtpRouteTypeEmail) {
	email.ID = o.ID
	email.RequestID = o.RequestID
	email.OtpID = o.OtpID
	email.Email = pgtype.Text{
		Valid:  true,
		String: o.Email,
	}
	email.CreatedAt = pgtype.Timestamptz{
		Time:  o.CreatedAt.Time,
		Valid: o.CreatedAt.Valid,
	}
	email.UpdatedAt = pgtype.Timestamptz{
		Time:  o.UpdatedAt.Time,
		Valid: o.UpdatedAt.Valid,
	}
	email.IsDeleted = pgtype.Bool{
		Valid: true,
		Bool:  o.IsDeleted,
	}
	email.DeletedAt = pgtype.Timestamptz{
		Time:  o.DeletedAt.Time,
		Valid: o.DeletedAt.Valid,
	}
	return
}
