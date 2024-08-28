package model

import (
	"fmt"
	"strings"
)

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
