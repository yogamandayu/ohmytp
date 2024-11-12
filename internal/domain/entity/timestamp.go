package entity

import "database/sql"

// Timestamp is entity timestamp.
type Timestamp struct {
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	IsDeleted bool         `json:"is_deleted"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
