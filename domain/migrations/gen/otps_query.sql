-- name: SaveOtp :one
INSERT INTO public.otps (id, request_id, route_type, code,
                         purpose, requested_at, confirmed_at, expired_at,
                         attempt, last_attempt_at,
                         resend_attempt, resend_at,
                         ip_address, user_agent,
                         created_at, updated_at)
VALUES ($1, $2, $3, $4,
        $5, $6, $7,
        $8, $9, $10, $11,
        $12, $13, $14,
        NOW(), NOW()) RETURNING *;

-- name: FindOtpByID :one
SELECT id,
       request_id,
       route_type,
       code,
       purpose,
       requested_at,
       confirmed_at,
       expired_at,
       attempt,
       last_attempt_at,
       resend_attempt,
       resend_at,
       ip_address,
       user_agent,
       created_at,
       updated_at
FROM public.otps
WHERE id = $1;

-- name: FindOtpByRequestID :one
SELECT id,
       request_id,
       route_type,
       code,
       purpose,
       requested_at,
       confirmed_at,
       expired_at,
       attempt,
       last_attempt_at,
       resend_attempt,
       resend_at,
       ip_address,
       user_agent,
       created_at,
       updated_at
FROM public.otps
WHERE request_id = $1;

-- name: UpdateOtp :one
UPDATE public.otps
SET request_id      = $2,
    route_type      = $3,
    code            = $4,
    purpose         = $5,
    requested_at    = $6,
    confirmed_at    = $7,
    expired_at      = $8,
    attempt         = $9,
    last_attempt_at = $10,
    resend_attempt  = $11,
    resend_at       = $12,
    ip_address      = $13,
    user_agent      = $14,
    updated_at      = NOW()
WHERE id = $1 RETURNING id, request_id, route_type, code,
                         purpose, requested_at, confirmed_at, expired_at,
                         attempt, last_attempt_at,
                         resend_attempt, resend_at,
                         ip_address, user_agent,
                         created_at, updated_at;
