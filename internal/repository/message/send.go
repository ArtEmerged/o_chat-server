package chat

import (
	"context"
	"fmt"

	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/ArtEmerged/o_chat-server/internal/repository/message/adapter"
)

// SendMessage sends message to chat by chat id and from user id.
func (r *messageRepo) SendMessage(ctx context.Context, in *model.SendMessageRequest) error {
	q := db.Query{
		Name: "message_repository.SendMessage",
	}

	q.QueryRaw =
		`WITH chat_exists AS (
		SELECT 1
		FROM public.chats
		WHERE id = $1 AND deleted_at IS NULL
	), user_in_chat AS (
		SELECT 1
		FROM public.chat_users
		WHERE chat_id = $1 AND user_id = $2
	)
	INSERT INTO public.chat_messages (chat_id, from_user_id, text, created_at)
	SELECT $1, $2, $3, $4
	WHERE EXISTS (SELECT 1 FROM chat_exists)
		AND EXISTS (SELECT 1 FROM user_in_chat);`

	msg := adapter.SendMessageRequestToRepo(in)

	result, err := r.db.DB().ExecContext(ctx, q, msg.ChatID, msg.FromUserID, msg.Text, msg.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed insert chat message: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat or user %w", model.ErrNotFound)
	}

	return nil
}
