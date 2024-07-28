-- +goose Up
CREATE TABLE chats (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    owner integer NOT NULL,
    created_at timestamptz  NOT NULL,
    deleted_at timestamptz
);

-- +goose Down
DROP TABLE chats;