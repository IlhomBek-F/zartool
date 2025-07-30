CREATE TABLE IF NOT EXISTS rent_tools (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    name TEXT,
    size TEXT,
    quantity BIGINT,
    user_id BIGINT,
    FOREIGN KEY (user_id) REFERENCES users(id)
)