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

	q := db.Query{
		Name: "chat_repository.CreateChat",
	}

	q.QueryRaw =
		`INSERT INTO public.chats (name, owner, created_at)
		VALUES ($1, $2, $3) 
		RETURNING id;`

	var chatID int64

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
		`INSERT INTO public.chat_users (chat_id, user_id) 
		VALUES %s`, values.String())

	_, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to add users to chat: %w", err)
	}

	return nil
}
