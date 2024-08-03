-- +goose Up
CREATE TABLE chats (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    owner integer NOT NULL,
    created_at timestamp  NOT NULL,
    deleted_at timestamp
);

-- +goose Down
DROP TABLE chats;