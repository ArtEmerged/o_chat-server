package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/ArtEmerged/o_chat-server/internal/client/db"
)

// DeleteChat deletes chat by id.
func (r *chatRepo) DeleteChat(ctx context.Context, id int64) error {
	query := fmt.Sprintf(
		`UPDATE %[1]s
		SET %[2]s = $1
		WHERE %[3]s = $2 AND %[2]s IS NULL;`,
		tableChatUsers,

		tableChatsDeletedAtColumn,
		tableChatsIDColumn,
	)

	q := db.Query{
		Name:     "deleted_repository.DeleteChat",
		QueryRaw: query,
	}

	createdAt := time.Now().UTC()

	_, err := r.db.DB().ExecContext(ctx, q, createdAt, id)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}
