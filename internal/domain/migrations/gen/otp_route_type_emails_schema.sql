CREATE TABLE IF NOT EXISTS public.otp_route_type_emails (
    id varchar(36) NOT NULL PRIMARY KEY,
    row_id SERIAL UNIQUE,
    otp_id varchar(36) NOT NULL ,
    request_id varchar(36) NOT NULL ,
    email varchar(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_deleted boolean,
    deleted_at timestamp with time zone
);
