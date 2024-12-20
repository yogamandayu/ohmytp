-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS public.otp_route_type_sms (
    id varchar(36) NOT NULL PRIMARY KEY,
    row_id SERIAL UNIQUE,
    otp_id varchar(36) NOT NULL ,
    request_id varchar(36) NOT NULL ,
    phone varchar(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_deleted boolean,
    deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_otp_route_type_sms_id ON public.otp_route_type_sms USING btree (id);
CREATE INDEX IF NOT EXISTS idx_otp_route_type_sms_otp_id ON public.otp_route_type_sms USING hash (otp_id);
CREATE INDEX IF NOT EXISTS idx_otp_route_type_sms_request_id ON public.otp_route_type_sms USING hash (request_id);
CREATE INDEX IF NOT EXISTS idx_otp_route_type_sms_row_id ON public.otp_route_type_sms USING hash (row_id);

---- create above / drop below ----

DROP TABLE otp_route_sms;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
