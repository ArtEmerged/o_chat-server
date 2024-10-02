package chat

import (
	"context"
	"log"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

// DeleteChat deletes chat by id.
func (s *chatService) DeleteChat(ctx context.Context, id int64) error {
	err := s.repo.DeleteChat(ctx, id)
	if err != nil {
		return err
	}

	err = s.cache.Del(ctx, model.ChatCacheKey(id))
	if err != nil {
		log.Printf("WARN: failed to delete chat in cache: %s\n", err.Error())
	}

	return nil
}
