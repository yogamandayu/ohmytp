-- name: SaveOtp :one
INSERT INTO public.otps (id, request_id, identifier, route_type, code,
                         purpose, requested_at, confirmed_at, expired_at,
                         attempt, last_attempt_at,
                         resend_attempt, resend_at,
                         ip_address, user_agent,
                         created_at, updated_at)
VALUES ($1, $2, $3, $4,
        $5, $6, $7,
        $8, $9, $10, $11,
        $12, $13, $14, $15,
        NOW(), NOW()) RETURNING *;

-- name: FindOtp :one
SELECT id,
       request_id,
       identifier,
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

-- name: FindOtpByIdentifierAndPurpose :one
SELECT id,
       request_id,
       identifier,
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
WHERE identifier = $1
  AND purpose = $2;

-- name: UpdateOtp :one
UPDATE public.otps
SET request_id      = $2,
    identifier      = $3,
    route_type      = $4,
    code            = $5,
    purpose         = $6,
    requested_at    = $7,
    confirmed_at    = $8,
    expired_at      = $9,
    attempt         = $10,
    last_attempt_at = $11,
    resend_attempt  = $12,
    resend_at       = $13,
    ip_address      = $14,
    user_agent      = $15,
    updated_at      = NOW()
WHERE id = $1 RETURNING id, request_id, identifier, route_type, code,
                         purpose, requested_at, confirmed_at, expired_at,
                         attempt, last_attempt_at,
                         resend_attempt, resend_at,
                         ip_address, user_agent,
                         created_at, updated_at;

-- name: UpdateOtpAttempt :one
UPDATE public.otps
SET attempt         = $2,
    last_attempt_at = $3,
    confirmed_at    = $4,
    updated_at      = NOW()
WHERE id = $1 RETURNING id, request_id, identifier, route_type, code,
                         purpose, requested_at, confirmed_at, expired_at,
                         attempt, last_attempt_at,
                         resend_attempt, resend_at,
                         ip_address, user_agent,
                         created_at, updated_at;
