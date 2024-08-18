package model

import "time"

// CreateChatRequest is a request for method CreateChat.
type CreateChatRequest struct {
	Name     string    `db:"name"`
	Owner    int64     `db:"owner"`
	CreateAt time.Time `db:"create_at"`
}
