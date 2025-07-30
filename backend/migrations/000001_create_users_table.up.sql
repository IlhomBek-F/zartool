CREATE TABLE IF NOT EXISTS users (
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    deleted_at  TIMESTAMPTZ,
    full_name   TEXT,
    address     TEXT,
    pre_payment BIGINT,        
    active      BOOLEAN,
    phones      TEXT,
    date        TEXT
);
