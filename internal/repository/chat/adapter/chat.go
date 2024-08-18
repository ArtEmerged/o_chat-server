package adapter

import (
	"time"

	"github.com/ArtEmerged/o_chat-server/internal/model"
	modelRepo "github.com/ArtEmerged/o_chat-server/internal/repository/chat/model"
)

// CreateChatRequestToRepo converts CreateChatRequest proto to repository CreateChatRequest
func CreateChatRequestToRepo(in *model.CreateChatRequest) *modelRepo.CreateChatRequest {
	return &modelRepo.CreateChatRequest{
		Name:     in.ChatName,
		Owner:    in.CreatorID,
		CreateAt: time.Now().UTC(),
	}
}
