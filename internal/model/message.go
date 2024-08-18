package model

// SendMessageRequest is a request for method SendMessage.
type SendMessageRequest struct {
	ChatID     int64
	FromUserID int64
	Text       string
}
