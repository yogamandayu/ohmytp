package entity

import "database/sql"

// Timestamp is entity timestamp.
type Timestamp struct {
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
