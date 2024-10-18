CREATE TABLE IF NOT EXISTS public.otps (
    id varchar(36) NOT NULL PRIMARY KEY,
    request_id varchar(36) NOT NULL ,
    route_type varchar(20),
    code varchar(20),
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
