package model

// CreateChatRequest is a request for method CreateChat.
type CreateChatRequest struct {
	ChatName  string
	CreatorID int64
	UserIDs   []int64
}
