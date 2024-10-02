package model

import (
	"fmt"
	"strings"
	"time"
)

const (
	DefaultTTL = time.Hour * 1
)

// SendMessageRequest is a request for method SendMessage.
type SendMessageRequest struct {
	ChatID     int64
	FromUserID int64
	Text       string
}

// Validate validates SendMessageRequest
func (r *SendMessageRequest) Validate() error {
	var errsText []string

	if r.Text == "" {
		errsText = append(errsText, "field text is required")
	}

	if len(errsText) > 0 {
		return fmt.Errorf("%w: %s", ErrInvalidArgument, strings.Join(errsText, ", "))
	}

	return nil
}

// Message is a message
type Message struct {
	ID         int64     `json:"id"`
	ChatID     int64     `json:"chat_id"`
	FromUserID int64     `json:"from_user_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}

// CreateMessageKey returns key for chat messages
func CreateMessageKey(chatID int64) string {
	return fmt.Sprintf("chats:chat:%d:messages", chatID)
}
