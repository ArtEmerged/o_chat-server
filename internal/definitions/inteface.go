package definitions

import "context"

// ChatService is an interface for chat service.
type ChatService interface {
	// CreateChat creates new chat by chat name and creator id with user ids.
	CreateChat(ctx context.Context, in *CreateChatRequest) (id int64, err error)
	// DeleteChat deletes chat by id.
	DeleteChat(ctx context.Context, id int64) error
	// SendMessage sends message to chat by chat id and from user id.
	SendMessage(ctx context.Context, in *SendMessageRequest) error
}

// ChatRepo is an interface for chat repository.
type ChatRepo interface {
	// CreateChat creates new chat by chat name and creator id with user ids.
	CreateChat(ctx context.Context, in *CreateChatRequest) (id int64, err error)
	// DeleteChat deletes chat by id.
	DeleteChat(ctx context.Context, id int64) error
	// SendMessage sends message to chat by chat id and from user id.
	SendMessage(ctx context.Context, in *SendMessageRequest) error
}
