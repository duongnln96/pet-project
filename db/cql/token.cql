CREATE TABLE IF NOT EXISTS auth_token (
    id uuid,
    user_id uuid,
    device_id uuid,
    user_agent text,
   	jwt_token text,
    remote_ip text,
    status text,
    expired_date timestamp,
    created_date timestamp,
    updated_date timestamp,
    PRIMARY KEY (id, expired_date)
);

