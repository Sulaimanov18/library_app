-- +goose Up
CREATE TABLE IF NOT EXISTS readers (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  email VARCHAR NOT NULL UNIQUE,
  phone VARCHAR,
  address TEXT,
  created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS readers; 