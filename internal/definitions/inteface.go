package definitions

import "context"

type ChatService interface {
	CreateChat(ctx context.Context, in *CreateChatRequest) (id int64, err error)
	DeleteChat(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, in *SendMessageRequest) error
}

type ChatRepo interface {
	CreateChat(ctx context.Context, in *CreateChatRequest) (id int64, err error)
	DeleteChat(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, in *SendMessageRequest) error
}
