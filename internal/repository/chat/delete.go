package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/ArtEmerged/library/client/db"
)

// DeleteChat deletes chat by id.
func (r *chatRepo) DeleteChat(ctx context.Context, id int64) error {
	q := db.Query{
		Name: "deleted_repository.DeleteChat",
	}

	q.QueryRaw =
		`UPDATE public.chats
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL;`

	createdAt := time.Now().UTC()

	_, err := r.db.DB().ExecContext(ctx, q, createdAt, id)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}
