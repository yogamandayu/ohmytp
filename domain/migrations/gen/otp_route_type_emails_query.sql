
-- name: SaveOtpRouteTypeEmail :one
INSERT INTO public.otp_route_type_emails (
    id, otp_id, request_id, email,
    created_at, updated_at
)
VALUES (
           $1, $2, $3, $4,
           NOW(), NOW()
       )
    RETURNING *;

-- name: FindOtpRouteTypeEmail :one
SELECT * FROM public.otp_route_type_emails
WHERE id = $1;

-- name: GetOtpRouteTypeEmails :many
SELECT * FROM public.otp_route_type_emails
ORDER BY created_at desc;

-- name: UpdateOtpRouteTypeEmail :one
UPDATE public.otp_route_type_emails SET
    otp_id = $2, request_id = $3, email = $4,
    updated_at = NOW()
WHERE id = $1
    RETURNING *;
