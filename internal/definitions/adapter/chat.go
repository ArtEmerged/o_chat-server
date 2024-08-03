package adapter

import (
	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

// CreateChatRequestToLocal converts CreateChatRequest proto to local CreateChatRequest
func CreateChatRequestToLocal(in *desc.CreateChatRequest) *def.CreateChatRequest {
	return &def.CreateChatRequest{
		ChatName:  in.GetChatName(),
		CreatorID: in.GetCreatorId(),
		UserIDs:   in.GetUserIds(),
	}
}

// SendMessageRequestToLocal converts SendMessageRequest proto to local SendMessageRequest
func SendMessageRequestToLocal(in *desc.SendMessageRequest) *def.SendMessageRequest {
	return &def.SendMessageRequest{
		ChatID:     in.GetChatId(),
		FromUserID: in.GetFromUserId(),
		Text:       in.GetText(),
	}
}
