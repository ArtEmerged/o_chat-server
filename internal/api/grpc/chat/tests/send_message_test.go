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

	"github.com/ArtEmerged/o_chat-server/internal/api/grpc/chat"
	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/ArtEmerged/o_chat-server/internal/service"
	"github.com/ArtEmerged/o_chat-server/internal/service/mocks"
	desc "github.com/ArtEmerged/o_chat-server/pkg/chat_v1"
)

func TestSendMessage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type messageServiceMockFunc func(mc *minimock.Controller) service.MessageService

	var (
		ctx                = context.Background()
		mc                 = minimock.NewController(t)
		chatID             = int64(gofakeit.Number(1, 1000))
		negativeChatID     = int64(gofakeit.Number(-1000, 0))
		fromUserID         = int64(gofakeit.Number(1, 1000))
		negativeFromUserID = int64(gofakeit.Number(-1000, 0))
		messageText        = gofakeit.City()
		serviceError       = fmt.Errorf("service error")
		serviceRequest     = &model.SendMessageRequest{
			ChatID:     chatID,
			FromUserID: fromUserID,
			Text:       messageText,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *desc.SendMessageResponse
		err                error
		chatServiceMock    chatServiceMockFunc
		messageServiceMock messageServiceMockFunc
	}{
		{
			name: "success send message",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					ChatId:     chatID,
					FromUserId: fromUserID,
					Text:       messageText,
				},
			},
			want:            &desc.SendMessageResponse{},
			chatServiceMock: func(mc *minimock.Controller) service.ChatService { return nil },
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := mocks.NewMessageServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, serviceRequest).Return(nil)
				return mock
			},
		},
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					ChatId:     chatID,
					FromUserId: fromUserID,
					Text:       messageText,
				},
			},
			want:            nil,
			err:             status.Error(codes.Internal, serviceError.Error()),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService { return nil },
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := mocks.NewMessageServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, serviceRequest).Return(serviceError)
				return mock
			},
		},
		{
			name: "negative chat id",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					ChatId:     negativeChatID,
					FromUserId: fromUserID,
					Text:       messageText,
				},
			},
			want:               nil,
			err:                status.Error(codes.InvalidArgument, "negative id"),
			chatServiceMock:    func(mc *minimock.Controller) service.ChatService { return nil },
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
		{
			name: "negative from user id",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					ChatId:     chatID,
					FromUserId: negativeFromUserID,
					Text:       messageText,
				},
			},
			want:               nil,
			err:                status.Error(codes.InvalidArgument, "negative from_user_id"),
			chatServiceMock:    func(mc *minimock.Controller) service.ChatService { return nil },
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.chatServiceMock(mc)
			messageServiceMock := tt.messageServiceMock(mc)
			s := chat.NewImplementation(chatServiceMock, messageServiceMock)
			resp, err := s.SendMessage(ctx, tt.args.req)

			require.Equal(t, tt.want, resp)
			require.Equal(t, tt.err, err)
		})
	}
}
