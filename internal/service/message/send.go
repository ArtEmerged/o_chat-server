package message

import (
	"context"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

// SendMessage sends message to chat by chat id and from user id.
func (s *messageService) SendMessage(ctx context.Context, in *model.SendMessageRequest) error {
	if err := in.Validate(); err != nil {
		return err
	}

	return s.repo.SendMessage(ctx, in)
}
