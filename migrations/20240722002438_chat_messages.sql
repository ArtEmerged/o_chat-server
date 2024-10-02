-- +goose Up
CREATE TABLE chat_messages (
    id serial PRIMARY KEY,
    chat_id integer NOT NULL,
    from_user_id integer NOT NULL,
    text text NOT NULL,
    created_at timestamp  NOT NULL
);

-- +goose Down
DROP TABLE chat_messages;
