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

// SendMessage sends message to chat.
func (s *Implementation) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	if in.GetFromUserId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative from_user_id")
	}

	if in.GetChatId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	err := s.messageService.SendMessage(ctx, adapter.SendMessageRequestFromProto(in))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		if errors.Is(err, model.ErrInvalidArgument) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.SendMessageResponse{}, nil
}
