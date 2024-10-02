package adapter

import (
	"time"

	"github.com/ArtEmerged/o_chat-server/internal/model"

	modelRepo "github.com/ArtEmerged/o_chat-server/internal/repository/message/model"
)

// SendMessageRequestToRepo converts SendMessageRequest local to repo SendMessageRequest
func SendMessageRequestToRepo(in *model.SendMessageRequest) *modelRepo.SendMessageRequest {
	return &modelRepo.SendMessageRequest{
		ChatID:     in.ChatID,
		FromUserID: in.FromUserID,
		Text:       in.Text,
		CreatedAt:  time.Now().UTC(),
	}
}

func MessageToModel(in *modelRepo.Message) *model.Message {
	return &model.Message{
		ID:         in.ID,
		ChatID:     in.ChatID,
		FromUserID: in.FromUserID,
		Text:       in.Text,
		CreatedAt:  in.CreatedAt,
	}
}
