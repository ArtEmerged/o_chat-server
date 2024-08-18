package chat

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

// DeleteChat deletes chat by id.
func (s *Implementation) DeleteChat(ctx context.Context, in *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	err := s.chatService.DeleteChat(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
