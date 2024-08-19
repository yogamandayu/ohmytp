package entity

import (
	"database/sql"
)

type OTP struct {
	ID              string       `json:"id"`
	RequestID       string       `json:"request_id"`
	AccessType      string       `json:"access_type"`
	AccessID        string       `json:"access_id"`
	RouteID         string       `json:"route_id"`
	RouteType       string       `json:"route"`
	Code            string       `json:"code"`
	RequestedAt     sql.NullTime `json:"request_at"`
	ConfirmedAt     sql.NullTime `json:"confirmed_at"`
	ExpiredAt       sql.NullTime `json:"expired_at"`
	Attempt         uint8        `json:"attempt"`
	LastAttemptedAt sql.NullTime `json:"last_attempted_at"`
	ResendAttempt   uint8        `json:"resend_attempt"`
	ResendAt        sql.NullTime `json:"resend_at"`
	IPAddress       string       `json:"ip_address"`
	UserAgent       string       `json:"user_agent"`
	Timestamp
}
