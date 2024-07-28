-- +goose Up
CREATE TABLE chat_messages (
    chat_id integer NOT NULL,
    from_user_id integer NOT NULL,
    text text NOT NULL,
    created_at timestamptz  NOT NULL
);

-- +goose Down
DROP TABLE chat_messages;
