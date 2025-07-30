CREATE TABLE IF NOT EXISTS owners (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ, 
    deleted_at TIMESTAMPTZ,
    login TEXT,
    password TEXT
);