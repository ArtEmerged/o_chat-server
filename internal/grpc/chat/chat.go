package grpc_chat

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

func (s *chatServer) CreateChat(ctx context.Context, in *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	if in.GetChatName() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing chat name")
	}

	if in.GetCreatorId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative creator id")
	}

	id, err := s.service.CreateChat(ctx, def.AdaptedCreateChatRequestToLocal(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateChatResponse{Id: int64(id)}, nil
}

func (s *chatServer) DeleteChat(ctx context.Context, in *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	err := s.service.DeleteChat(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}

func (s *chatServer) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	if in.GetText() == "" || in.GetFromUserId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "missing field from or text")
	}

	if in.GetChatId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	err := s.service.SendMessage(ctx, def.AdaptedSendMessageRequestToLocal(in))
	if err != nil {
		if errors.Is(err, def.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.SendMessageResponse{}, nil
}
