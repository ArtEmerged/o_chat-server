package chat

import (
	"context"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

// SendMessage sends message to chat by chat id and from user id.
func (s *messageService) SendMessage(ctx context.Context, in *model.SendMessageRequest) error {
	return s.repo.SendMessage(ctx, in)
}
