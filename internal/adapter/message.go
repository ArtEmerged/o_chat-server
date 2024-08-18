package adapter

import (
	"github.com/ArtEmerged/o_chat-server/internal/model"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

// SendMessageRequestFromProto converts SendMessageRequest proto to local SendMessageRequest
func SendMessageRequestFromProto(in *desc.SendMessageRequest) *model.SendMessageRequest {
	return &model.SendMessageRequest{
		ChatID:     in.GetChatId(),
		FromUserID: in.GetFromUserId(),
		Text:       in.GetText(),
	}
}
