package repository

import (
	"context"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

// ChatRepo is an interface for chat repository.
type ChatRepo interface {
	// CreateChat creates new chat by chat name and creator id with user ids.
	CreateChat(ctx context.Context, in *model.CreateChatRequest) (id int64, err error)
	// DeleteChat deletes chat by id.
	DeleteChat(ctx context.Context, id int64) error
	// AddUsersToChat adds users to chat.
	AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error
}

// MessageRepo is an interface for chat repository.
type MessageRepo interface {
	// SendMessage sends message to chat by chat id and from user id.
	SendMessage(ctx context.Context, in *model.SendMessageRequest) (*model.Message, error)
}
