package chat

import (
	"github.com/ArtEmerged/o_chat-server/internal/service"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

var _ desc.ChatV1Server = (*Implementation)(nil)

// Implementation implements chat gRPC interface.
type Implementation struct {
	desc.UnimplementedChatV1Server

	chatService    service.ChatService
	messageService service.MessageService
}

// NewImplementation registers the chat service on the gRPC server.
// s - pointer to the gRPC server
// service - the chat service interface to be registered
func NewImplementation(chatService service.ChatService, messageService service.MessageService) *Implementation {
	return &Implementation{
		chatService:    chatService,
		messageService: messageService,
	}
}
