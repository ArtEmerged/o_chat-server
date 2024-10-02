package message

import (
	"context"
	"fmt"

	"github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/ArtEmerged/o_chat-server/internal/repository/message/adapter"
	repoModel "github.com/ArtEmerged/o_chat-server/internal/repository/message/model"
)

// SendMessage sends message to chat by chat id and from user id.
func (r *messageRepo) SendMessage(ctx context.Context, in *model.SendMessageRequest) (*model.Message, error) {
	q := db.Query{
		Name: "message_repository.SendMessage",
	}

	q.QueryRaw =`
	INSERT INTO public.chat_messages (chat_id, from_user_id, text, created_at)
		SELECT $1, $2, $3, $4
	WHERE EXISTS (
    	SELECT 1 
    	FROM public.chats 
    	WHERE id = $1 AND deleted_at IS NULL
	)
	AND EXISTS (
    	SELECT 1 
    	FROM public.chat_users 
    	WHERE chat_id = $1 AND user_id = $2
	)
	RETURNING id, chat_id, from_user_id, text, created_at;`

	msg := adapter.SendMessageRequestToRepo(in)

	message := new(repoModel.Message)
	err := r.db.DB().ScanOneContext(
		ctx,
		message,
		q,
		msg.ChatID,
		msg.FromUserID,
		msg.Text,
		msg.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed insert chat message: %w", err)
	}

	return adapter.MessageToModel(message), nil
}
