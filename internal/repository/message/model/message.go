package model

import "time"

// SendMessageRequest is a request for method SendMessage.
type SendMessageRequest struct {
	ChatID     int64     `db:"chat_id"`
	FromUserID int64     `db:"from_user_id"`
	Text       string    `db:"text"`
	CreatedAt  time.Time `db:"created_at"`
}

type Message struct {
	ID         int64     `db:"id"`
	ChatID     int64     `db:"chat_id"`
	FromUserID int64     `db:"from_user_id"`
	Text       string    `db:"text"`
	CreatedAt  time.Time `db:"created_at"`
}
