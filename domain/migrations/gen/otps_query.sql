
-- name: SaveOtp :one
INSERT INTO public.otps (
    id, request_id, route_type, code,
    requested_at, confirmed_at, expired_at,
    attempt, last_attempt_at,
    resend_attempt, resend_at,
    ip_address, user_agent,
    created_at, updated_at
)
VALUES (
           $1, $2, $3, $4,
           $5, $6, $7,
           $8, $9, $10, $11,
           $12, $13,
           $14, $15
       )
    RETURNING *;

-- name: FindOtp :one
SELECT * FROM public.otps
WHERE id = $1;

-- name: GetOtps :many
SELECT * FROM public.otps
ORDER BY created_at desc;