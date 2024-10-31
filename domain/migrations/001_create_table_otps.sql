-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS public.otps (
    id varchar(36) NOT NULL PRIMARY KEY,
    row_id SERIAL UNIQUE,
    request_id varchar(36) NOT NULL ,
    route_type varchar(20),
    code varchar(20),
    purpose varchar(100),
    requested_at timestamp with time zone,
    confirmed_at timestamp with time zone,
    expired_at timestamp with time zone,
    attempt smallint,
    last_attempt_at timestamp with time zone,
    resend_attempt smallint,
    resend_at timestamp with time zone,
    ip_address varchar(100),
    user_agent text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_deleted boolean,
    deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_otps_id ON public.otps USING btree (id);
CREATE INDEX IF NOT EXISTS idx_otps_request_id ON public.otps USING hash (request_id);

---- create above / drop below ----

DROP TABLE otps;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
