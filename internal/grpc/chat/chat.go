package grpc_chat

import (
	"context"

	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ desc.ChatV1Server = (*chatServer)(nil)

func (s *chatServer) CreateChat(ctx context.Context, in *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	if len(in.GetUserName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty user name list")
	}
	id := gofakeit.Number(1, 99999)
	return &desc.CreateChatResponse{Id: int64(id)}, nil
}
func (s *chatServer) DeleteChat(ctx context.Context, in *desc.DeleteChatRequest) (*desc.DeleteChatResponse, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}
	return nil, nil
}
func (s *chatServer) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	if in.GetText() == "" || in.GetFrom() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing field from or text")
	}
	return nil, nil
}
