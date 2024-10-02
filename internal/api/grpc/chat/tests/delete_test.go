package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ArtEmerged/o_chat-server/internal/api/grpc/chat"
	"github.com/ArtEmerged/o_chat-server/internal/service"
	"github.com/ArtEmerged/o_chat-server/internal/service/mocks"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

func TestDeleteChate(t *testing.T) {
	type args struct {
		ctx context.Context
		req *desc.DeleteChatRequest
	}

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type messageServiceMockFunc func(mc *minimock.Controller) service.MessageService

	var (
		ctx            = context.Background()
		mc             = minimock.NewController(t)
		chatID         = int64(gofakeit.Number(1, 1000))
		negativeChatID = int64(gofakeit.Number(-1000, 0))
		serviceError   = fmt.Errorf("service error")
	)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		chatServiceMock    chatServiceMockFunc
		messageServiceMock messageServiceMockFunc
	}{
		{
			name: "success delete chat",
			args: args{
				ctx: ctx,
				req: &desc.DeleteChatRequest{
					Id: chatID,
				},
			},
			want: nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteChatMock.Expect(ctx, chatID).Return(nil)
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
		{
			name: "negative chat id",
			args: args{
				ctx: ctx,
				req: &desc.DeleteChatRequest{
					Id: negativeChatID,
				},
			},
			want:               nil,
			err:                status.Error(codes.InvalidArgument, "negative id"),
			chatServiceMock:    func(mc *minimock.Controller) service.ChatService { return nil },
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: &desc.DeleteChatRequest{
					Id: chatID,
				},
			},
			want: nil,
			err:  status.Error(codes.Internal, serviceError.Error()),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteChatMock.Expect(ctx, chatID).Return(serviceError)
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.chatServiceMock(mc)
			messageServiceMock := tt.messageServiceMock(mc)
			s := chat.NewImplementation(chatServiceMock, messageServiceMock)
			resp, err := s.DeleteChat(ctx, tt.args.req)

			require.Equal(t, tt.want, resp)
			require.Equal(t, tt.err, err)
		})
	}
}
