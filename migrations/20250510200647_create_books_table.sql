-- +goose Up
CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  title VARCHAR NOT NULL,
  description TEXT,
  publisher VARCHAR,
  count INT DEFAULT 1,
  created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS books;
