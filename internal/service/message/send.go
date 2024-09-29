package message

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

// SendMessage sends message to chat by chat id and from user id.
func (s *messageService) SendMessage(ctx context.Context, in *model.SendMessageRequest) error {
	if err := in.Validate(); err != nil {
		return err
	}

	message, err := s.repo.SendMessage(ctx, in)
	if err != nil {
		return err
	}

	msgIDStr := strconv.FormatInt(message.ID, 10)

	err = s.cache.HSet(ctx, model.CreateMessageKey(in.ChatID), msgIDStr, message, time.Hour*1)
	if err != nil {
		log.Printf("WARN: failed to save message in cache: %s\n", err.Error())
	}

	return nil
}
