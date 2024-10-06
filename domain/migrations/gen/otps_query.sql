
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
           NOW(), NOW()
       )
    RETURNING *;

-- name: FindOtp :one
SELECT * FROM public.otps
WHERE id = $1;

-- name: GetOtps :many
SELECT * FROM public.otps
ORDER BY created_at desc;

-- name: UpdateOtp :one
UPDATE public.otps SET
    request_id = $2, route_type = $3, code = $4,
    requested_at = $5, confirmed_at = $6, expired_at = $7,
    attempt = $8, last_attempt_at = $9,
    resend_attempt = $10, resend_at = $11,
    ip_address = $12, user_agent = $13,
    updated_at = NOW()
WHERE id = $1
    RETURNING *;
