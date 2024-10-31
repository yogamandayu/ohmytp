package entity

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yogamandayu/ohmytp/storage/repository"
)

// Otp is a struct to hold all data related to otp action.
type Otp struct {
	ID            string       `json:"id"`
	RowID         int          `json:"row_id"`
	RequestID     string       `json:"request_id"`
	RouteType     string       `json:"route"`
	Code          string       `json:"code"`
	Purpose       string       `json:"purpose"`
	RequestedAt   sql.NullTime `json:"request_at"`
	ConfirmedAt   sql.NullTime `json:"confirmed_at"`
	ExpiredAt     sql.NullTime `json:"expired_at"`
	Attempt       uint8        `json:"attempt"`
	LastAttemptAt sql.NullTime `json:"last_attempt_at"`
	ResendAttempt uint8        `json:"resend_attempt"`
	ResendAt      sql.NullTime `json:"resend_at"`
	IPAddress     string       `json:"ip_address"`
	UserAgent     string       `json:"user_agent"`
	Timestamp
}

// SetWithOtpRepository is to set entity otp with repository otp.
func (o *Otp) SetWithOtpRepository(otp repository.Otp) {
	o.ID = otp.ID
	o.RequestID = otp.RequestID
	o.RouteType = otp.RouteType.String
	o.Code = otp.Code.String
	o.RequestedAt = sql.NullTime{
		Time:  otp.RequestedAt.Time,
		Valid: otp.RequestedAt.Valid,
	}
	o.ConfirmedAt = sql.NullTime{
		Time:  otp.ConfirmedAt.Time,
		Valid: otp.ConfirmedAt.Valid,
	}
	o.ExpiredAt = sql.NullTime{
		Time:  otp.ExpiredAt.Time,
		Valid: otp.ExpiredAt.Valid,
	}
	o.Attempt = uint8(otp.Attempt.Int16)
	o.LastAttemptAt = sql.NullTime{
		Time:  otp.LastAttemptAt.Time,
		Valid: otp.LastAttemptAt.Valid,
	}
	o.ResendAttempt = uint8(otp.ResendAttempt.Int16)
	o.ResendAt = sql.NullTime{
		Time:  otp.ResendAt.Time,
		Valid: otp.ResendAt.Valid,
	}
	o.IPAddress = otp.IpAddress.String
	o.UserAgent = otp.UserAgent.String
	o.CreatedAt = sql.NullTime{
		Time:  otp.CreatedAt.Time,
		Valid: otp.CreatedAt.Valid,
	}
	o.UpdatedAt = sql.NullTime{
		Time:  otp.UpdatedAt.Time,
		Valid: otp.UpdatedAt.Valid,
	}
	o.IsDeleted = otp.IsDeleted.Bool && otp.IsDeleted.Valid
	o.DeletedAt = sql.NullTime{
		Time:  otp.DeletedAt.Time,
		Valid: otp.DeletedAt.Valid,
	}
}

// TransformToOtpRepository is to transform entity otp to repository otp.
func (o *Otp) TransformToOtpRepository() (otp repository.Otp) {
	otp.ID = o.ID
	otp.RequestID = o.RequestID
	otp.RouteType = pgtype.Text{
		Valid:  true,
		String: o.RouteType,
	}
	otp.Code = pgtype.Text{
		Valid:  true,
		String: o.Code,
	}
	otp.RequestedAt = pgtype.Timestamptz{
		Time:  o.RequestedAt.Time,
		Valid: o.RequestedAt.Valid,
	}
	otp.ConfirmedAt = pgtype.Timestamptz{
		Time:  o.ConfirmedAt.Time,
		Valid: o.ConfirmedAt.Valid,
	}
	otp.ExpiredAt = pgtype.Timestamptz{
		Time:  o.ExpiredAt.Time,
		Valid: o.ExpiredAt.Valid,
	}
	otp.Attempt.Int16 = int16(o.Attempt)
	otp.LastAttemptAt = pgtype.Timestamptz{
		Time:  o.LastAttemptAt.Time,
		Valid: o.LastAttemptAt.Valid,
	}
	otp.ResendAttempt.Int16 = int16(o.ResendAttempt)
	otp.ResendAt = pgtype.Timestamptz{
		Time:  o.ResendAt.Time,
		Valid: o.ResendAt.Valid,
	}
	otp.IpAddress = pgtype.Text{
		Valid:  true,
		String: o.IPAddress,
	}
	otp.UserAgent = pgtype.Text{
		Valid:  true,
		String: o.UserAgent,
	}
	otp.CreatedAt = pgtype.Timestamptz{
		Time:  o.CreatedAt.Time,
		Valid: o.CreatedAt.Valid,
	}
	otp.UpdatedAt = pgtype.Timestamptz{
		Time:  o.UpdatedAt.Time,
		Valid: o.UpdatedAt.Valid,
	}
	otp.DeletedAt = pgtype.Timestamptz{
		Time:  o.DeletedAt.Time,
		Valid: o.DeletedAt.Valid,
	}
	return
}
