-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price_in_cents INT NOT NULL CHECK (price_in_cents >= 0),
    quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
SELECT 'down SQL query';
