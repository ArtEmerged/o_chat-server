package grpc_chat

import (
	"google.golang.org/grpc"

	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

var _ desc.ChatV1Server = (*chatServer)(nil)

type chatServer struct {
	desc.UnimplementedChatV1Server
	service def.ChatService
}

func Register(s *grpc.Server, service def.ChatService) {
	desc.RegisterChatV1Server(s, &chatServer{service: service})
}
