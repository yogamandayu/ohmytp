// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: otp_route_type_emails_query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const findOtpRouteTypeEmail = `-- name: FindOtpRouteTypeEmail :one
SELECT id, row_id, otp_id, request_id, email, created_at, updated_at, is_deleted, deleted_at
FROM public.otp_route_type_emails
WHERE id = $1
`

func (q *Queries) FindOtpRouteTypeEmail(ctx context.Context, id string) (OtpRouteTypeEmail, error) {
	row := q.db.QueryRow(ctx, findOtpRouteTypeEmail, id)
	var i OtpRouteTypeEmail
	err := row.Scan(
		&i.ID,
		&i.RowID,
		&i.OtpID,
		&i.RequestID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.DeletedAt,
	)
	return i, err
}

const getOtpRouteTypeEmails = `-- name: GetOtpRouteTypeEmails :many
SELECT id, row_id, otp_id, request_id, email, created_at, updated_at, is_deleted, deleted_at
FROM public.otp_route_type_emails
ORDER BY created_at desc
`

func (q *Queries) GetOtpRouteTypeEmails(ctx context.Context) ([]OtpRouteTypeEmail, error) {
	rows, err := q.db.Query(ctx, getOtpRouteTypeEmails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OtpRouteTypeEmail
	for rows.Next() {
		var i OtpRouteTypeEmail
		if err := rows.Scan(
			&i.ID,
			&i.RowID,
			&i.OtpID,
			&i.RequestID,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveOtpRouteTypeEmail = `-- name: SaveOtpRouteTypeEmail :one
INSERT INTO public.otp_route_type_emails (id, otp_id, request_id, email,
                                          created_at, updated_at)
VALUES ($1, $2, $3, $4,
        NOW(), NOW()) RETURNING id, row_id, otp_id, request_id, email, created_at, updated_at, is_deleted, deleted_at
`

type SaveOtpRouteTypeEmailParams struct {
	ID        string
	OtpID     string
	RequestID string
	Email     pgtype.Text
}

func (q *Queries) SaveOtpRouteTypeEmail(ctx context.Context, arg SaveOtpRouteTypeEmailParams) (OtpRouteTypeEmail, error) {
	row := q.db.QueryRow(ctx, saveOtpRouteTypeEmail,
		arg.ID,
		arg.OtpID,
		arg.RequestID,
		arg.Email,
	)
	var i OtpRouteTypeEmail
	err := row.Scan(
		&i.ID,
		&i.RowID,
		&i.OtpID,
		&i.RequestID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.DeletedAt,
	)
	return i, err
}

const updateOtpRouteTypeEmail = `-- name: UpdateOtpRouteTypeEmail :one
UPDATE public.otp_route_type_emails
SET otp_id     = $2,
    request_id = $3,
    email      = $4,
    updated_at = NOW()
WHERE id = $1 RETURNING id, row_id, otp_id, request_id, email, created_at, updated_at, is_deleted, deleted_at
`

type UpdateOtpRouteTypeEmailParams struct {
	ID        string
	OtpID     string
	RequestID string
	Email     pgtype.Text
}

func (q *Queries) UpdateOtpRouteTypeEmail(ctx context.Context, arg UpdateOtpRouteTypeEmailParams) (OtpRouteTypeEmail, error) {
	row := q.db.QueryRow(ctx, updateOtpRouteTypeEmail,
		arg.ID,
		arg.OtpID,
		arg.RequestID,
		arg.Email,
	)
	var i OtpRouteTypeEmail
	err := row.Scan(
		&i.ID,
		&i.RowID,
		&i.OtpID,
		&i.RequestID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.DeletedAt,
	)
	return i, err
}
