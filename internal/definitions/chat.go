package definitions

// CreateChatRequest is a request for method CreateChat.
type CreateChatRequest struct {
	ChatName  string
	CreatorID int64
	UserIDs   []int64
}

// SendMessageRequest is a request for method SendMessage.
type SendMessageRequest struct {
	ChatID     int64
	FromUserID int64
	Text       string
}
