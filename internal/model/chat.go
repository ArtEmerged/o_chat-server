package model

import (
	"fmt"
	"strings"
)

const chatCacheKey = "chats:chat:%d"

// ChatCacheKey returns key for chat
func ChatCacheKey(chatID int64) string {
	return fmt.Sprintf(chatCacheKey, chatID)
}

// Chat is a model for chat.
type Chat struct {
	ID        int64   `json:"id"`
	ChatName  string  `json:"chat_name"`
	CreatorID int64   `json:"creator_id"`
	UserIDs   []int64 `json:"user_ids"`
}

// CreateChatRequest is a request for method CreateChat.
type CreateChatRequest struct {
	ChatName  string
	CreatorID int64
	UserIDs   []int64
}

// Validate validates CreateChatRequest.
func (r *CreateChatRequest) Validate() error {
	var errsText []string

	// validate required fields
	if r.ChatName == "" {
		errsText = append(errsText, "field chat_name is required")
	}

	if len(errsText) > 0 {
		return fmt.Errorf("%w: %s", ErrInvalidArgument, strings.Join(errsText, ", "))
	}

	return nil
}
