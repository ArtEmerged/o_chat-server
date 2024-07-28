package definitions

import (
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

type CreateChatRequest struct {
	ChatName  string
	CreatorID int64
	UserIDs   []int64
}

type SendMessageRequest struct {
	ChatID     int64
	FromUserID int64
	Text       string
}

func AdaptedCreateChatRequestToLocal(in *desc.CreateChatRequest) *CreateChatRequest {
	return &CreateChatRequest{
		ChatName:  in.GetChatName(),
		CreatorID: in.GetCreatorId(),
		UserIDs:   in.GetUserIds(),
	}
}

func AdaptedSendMessageRequestToLocal(in *desc.SendMessageRequest) *SendMessageRequest {
	return &SendMessageRequest{
		ChatID:     in.GetChatId(),
		FromUserID: in.GetFromUserId(),
		Text:       in.GetText(),
	}
}
