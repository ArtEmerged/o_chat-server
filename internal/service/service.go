package service

import (
	"context"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

// ChatService is an interface for chat service.
type ChatService interface {
	// CreateChat creates new chat by chat name and creator id with user ids.
	CreateChat(ctx context.Context, in *model.CreateChatRequest) (id int64, err error)
	// DeleteChat deletes chat by id.
	DeleteChat(ctx context.Context, id int64) error
}

// MessageService is an interface for message service.
type MessageService interface {
	// SendMessage sends message to chat by chat id and from user id.
	SendMessage(ctx context.Context, in *model.SendMessageRequest) error
}
