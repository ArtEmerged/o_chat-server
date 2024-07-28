-- +goose Up


CREATE TABLE chat_users (
    chat_id integer NOT NULL,
    user_id integer NOT NULL
);

-- +goose Down
DROP TABLE chat_users;
