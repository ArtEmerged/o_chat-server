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
	query := fmt.Sprintf(
		`WITH chat_exists AS (
		SELECT 1
		FROM %[1]s
		WHERE %[2]s = $1 AND %[3]s IS NULL
	), user_in_chat AS (
		SELECT 1
		FROM %[4]s
		WHERE %[5]s = $1 AND %[6]s = $2
	)
	INSERT INTO %[7]s (%[8]s, %[9]s, %[10]s, %[11]s)
	SELECT $1, $2, $3, $4
	WHERE EXISTS (SELECT 1 FROM chat_exists)
	  AND EXISTS (SELECT 1 FROM user_in_chat);`,
		tableChats,                // 1
		tableChatsIDColumn,        // 2
		tableChatsDeletedAtColumn, // 3

		tableChatUsers,             // 4
		tableChatUsersChatIDColumn, // 5
		tableChatUsersUserIDColumn, // 6

		tableMessages,                 // 7
		tableMessagesChatIDColumn,     // 8
		tableMessagesFromUserIDColumn, // 9
		tableMessagesTextColumn,       // 10
		tableMessagesCreatedAtColumn,  // 11
	)

	msg := adapter.SendMessageRequestToRepo(in)

	q := db.Query{
		Name:     "message_repository.SendMessage",
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, msg.ChatID, msg.FromUserID, msg.Text, msg.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed insert chat message: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat or user %w", model.ErrNotFound)
	}

	return nil
}
