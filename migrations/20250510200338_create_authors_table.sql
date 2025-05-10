-- +goose Up
CREATE TABLE IF NOT EXISTS authors (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  bio TEXT,
  created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS authors;
