package chat

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ArtEmerged/o_chat-server/internal/adapter"
	"github.com/ArtEmerged/o_chat-server/internal/model"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

// CreateChat creates new chat by chat name and creator id with user ids.
func (s *Implementation) CreateChat(ctx context.Context, in *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	if in.GetCreatorId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative creator id")
	}

	id, err := s.chatService.CreateChat(ctx, adapter.CreateChatRequestFromProto(in))
	if err != nil {
		if errors.Is(err, model.ErrInvalidArgument) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateChatResponse{Id: int64(id)}, nil
}
