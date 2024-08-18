package chat

import (
	"context"
	"fmt"
	"strings"

	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/ArtEmerged/o_chat-server/internal/repository/chat/adapter"
)

// CreateChat creates new chat by chat name and creator id.
func (r *chatRepo) CreateChat(ctx context.Context, in *model.CreateChatRequest) (id int64, err error) {
	newChat := adapter.CreateChatRequestToRepo(in)

	query := fmt.Sprintf(
		`INSERT INTO	%[1]s (%[2]s, %[3]s, %[4]s)
		VALUES ($1, $2, $3) RETURNING %[5]s;`,
		tableChats,

		tableChatsNameColumn,
		tableChatsOwnerColumn,
		tableChatsCreatedAtColumn,
		tableChatsIDColumn,
	)

	var chatID int64

	q := db.Query{
		Name:     "chat_repository.CreateChat",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, newChat.Name, newChat.Owner, newChat.CreateAt).Scan(&chatID)
	if err != nil {
		return -1, fmt.Errorf("failed to create chat: %w", err)
	}

	return chatID, nil
}

// AddUsersToChat adds users to chat.
func (r *chatRepo) AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error {
	args := []interface{}{chatID}
	values := strings.Builder{}

	for i, userID := range userIDs {
		if i > 0 {
			values.WriteString(", ")
		}
		values.WriteString(fmt.Sprintf("($1, $%d)", i+2))
		args = append(args, userID)

	}

	q := db.Query{
		Name: "chat_repository.AddUsersToChat",
	}

	q.QueryRaw = fmt.Sprintf(
		`INSERT INTO %[1]s (%[2]s, %[3]s)
		VALUES %[4]s`,
		tableChatUsers,

		tableChatUsersChatIDColumn,
		tableChatUsersUserIDColumn,
		values.String(),
	)

	_, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to add users to chat: %w", err)
	}

	return nil
}
