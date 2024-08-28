package adapter

import (
	"github.com/ArtEmerged/o_chat-server/internal/model"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

// CreateChatRequestFromProto converts CreateChatRequest proto to local CreateChatRequest
func CreateChatRequestFromProto(in *desc.CreateChatRequest) *model.CreateChatRequest {
	return &model.CreateChatRequest{
		ChatName:  in.GetChatName(),
		CreatorID: in.GetCreatorId(),
		UserIDs:   in.GetUserIds(),
	}
}
