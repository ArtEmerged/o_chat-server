package chat

import (
	"context"
)

// DeleteChat deletes chat by id.
func (s *chatService) DeleteChat(ctx context.Context, id int64) error {
	return s.repo.DeleteChat(ctx, id)
}
