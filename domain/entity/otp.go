package entity

import (
	"database/sql"
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
