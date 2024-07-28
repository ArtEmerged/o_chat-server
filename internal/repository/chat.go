package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
)

func (r *chatRepo) CreateChat(ctx context.Context, in *def.CreateChatRequest) (id int64, err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return -1, fmt.Errorf("failed to begin transaction: %w", err)
	}

	q := `
	INSERT INTO	public.chats (name, owner, created_at)
	VALUES ($1, $2, $3) RETURNING id;`
	var chatID int64

	err = tx.QueryRow(ctx, q, in.ChatName, in.CreatorID, time.Now()).Scan(&chatID)
	if err != nil {
		tx.Rollback(ctx)
		return -1, fmt.Errorf("failed to create chat: %w", err)
	}

	args := []interface{}{chatID}
	values := strings.Builder{}

	for i, user_id := range in.UserIDs {
		if i > 0 {
			values.WriteString(", ")
		}
		values.WriteString(fmt.Sprintf("($1, $%d)", i+2))
		args = append(args, user_id)

	}

	q = fmt.Sprintf(`
	INSERT INTO public.chat_users (chat_id, user_id)
	VALUES %s`, values.String())

	_, err = tx.Exec(ctx, q, args...)
	if err != nil {
		tx.Rollback(ctx)
		return -1, fmt.Errorf("failed to add users to chat: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return -1, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return chatID, nil
}

func (r *chatRepo) DeleteChat(ctx context.Context, id int64) error {
	q := `
	UPDATE public.chats
	SET deleted_at = $1
	WHERE id = $2 AND deleted_at IS NULL;`

	_, err := r.db.Exec(ctx, q, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}
func (r *chatRepo) SendMessage(ctx context.Context, in *def.SendMessageRequest) error {
	q := `
	WITH chat_exists AS (
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

	result, err := r.db.Exec(ctx, q, in.ChatID, in.FromUserID, in.Text, time.Now())
	if err != nil {
		return fmt.Errorf("failed insert chat message: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat or user %w", def.ErrNotFound)
	}

	return nil
}
