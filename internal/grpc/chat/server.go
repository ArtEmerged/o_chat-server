package grpc_chat

import (
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
)

type chatServer struct {
	desc.UnimplementedChatV1Server
}

func Register(s *grpc.Server) {
	desc.RegisterChatV1Server(s, &chatServer{})
}
