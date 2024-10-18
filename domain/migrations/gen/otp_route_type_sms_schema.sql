CREATE TABLE IF NOT EXISTS public.otp_route_type_sms (
    id varchar(36) NOT NULL PRIMARY KEY,
    otp_id varchar(36) NOT NULL ,
    request_id varchar(36) NOT NULL ,
    phone varchar(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_deleted boolean,
    deleted_at timestamp with time zone
);
