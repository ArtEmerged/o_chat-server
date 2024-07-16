package grpc_chat

import (
	"google.golang.org/grpc"

	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

type chatServer struct {
	desc.UnimplementedChatV1Server
}

func Register(s *grpc.Server) {
	desc.RegisterChatV1Server(s, &chatServer{})
}
