-- name: SaveOtpRouteTypeSMS :one
INSERT INTO public.otp_route_type_sms (id, otp_id, request_id, phone,
                                       created_at, updated_at)
VALUES ($1, $2, $3, $4,
        NOW(), NOW()) RETURNING *;

-- name: FindOtpRouteTypeSMS :one
SELECT *
FROM public.otp_route_type_sms
WHERE id = $1;

-- name: GetOtpRouteTypeSMS :many
SELECT *
FROM public.otp_route_type_sms
ORDER BY created_at desc;

-- name: UpdateOtpRouteTypeSMS :one
UPDATE public.otp_route_type_sms
SET otp_id     = $2,
    request_id = $3,
    phone      = $4,
    updated_at = NOW()
WHERE id = $1 RETURNING *;
