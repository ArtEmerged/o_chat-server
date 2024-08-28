package model

import (
	"fmt"
	"strings"
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
