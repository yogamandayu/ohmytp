-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS public.otp_route_type_emails (
    id varchar(36) NOT NULL PRIMARY KEY,
    otp_id varchar(36) NOT NULL ,
    request_id varchar(36) NOT NULL ,
    email varchar(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_deleted boolean,
    deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_otp_route_type_emails_id ON public.otp_route_type_emails USING btree (id);
CREATE INDEX IF NOT EXISTS idx_otp_route_type_emails_otp_id ON public.otp_route_type_emails USING hash (otp_id);
CREATE INDEX IF NOT EXISTS idx_otp_route_type_emails_request_id ON public.otp_route_type_emails USING hash (request_id);

---- create above / drop below ----

DROP TABLE otps

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
