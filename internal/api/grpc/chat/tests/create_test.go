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

func TestCreateChat(t *testing.T) {

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type messageServiceMockFunc func(mc *minimock.Controller) service.MessageService

	type args struct {
		ctx context.Context
		req *desc.CreateChatRequest
	}

	var (
		mc                = minimock.NewController(t)
		ctx               = context.Background()
		chatName          = gofakeit.StreetName()
		chatID            = gofakeit.Int64()
		creatorID         = int64(gofakeit.Number(1, 100))
		creatorIDNegative = int64(gofakeit.Number(-100, 0))
		userIDs           = []int64{
			gofakeit.Int64(),
			gofakeit.Int64(),
			gofakeit.Int64(),
			gofakeit.Int64(),
		}
		serviceRequest = &model.CreateChatRequest{
			ChatName:  chatName,
			CreatorID: creatorID,
			UserIDs:   userIDs,
		}
		serviceError = fmt.Errorf("service error")
	)

	tests := []struct {
		name               string
		args               args
		want               *desc.CreateChatResponse
		err                error
		chatServiceMock    chatServiceMockFunc
		messageServiceMock messageServiceMockFunc
	}{
		{
			name: "success create chat",
			args: args{
				ctx: ctx,
				req: &desc.CreateChatRequest{
					ChatName:  chatName,
					CreatorId: creatorID,
					UserIds:   userIDs,
				},
			},
			want: &desc.CreateChatResponse{Id: chatID},
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, serviceRequest).Return(chatID, nil)

				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: &desc.CreateChatRequest{
					ChatName:  chatName,
					CreatorId: creatorID,
					UserIds:   userIDs,
				},
			},
			want: nil,
			err:  status.Error(codes.Internal, serviceError.Error()),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, serviceRequest).Return(-1, serviceError)
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
		{
			name: "negative creator id",
			args: args{
				ctx: ctx,
				req: &desc.CreateChatRequest{
					ChatName:  chatName,
					CreatorId: creatorIDNegative,
					UserIds:   userIDs,
				},
			},
			want:               nil,
			err:                status.Error(codes.InvalidArgument, "negative creator id"),
			chatServiceMock:    func(mc *minimock.Controller) service.ChatService { return nil },
			messageServiceMock: func(mc *minimock.Controller) service.MessageService { return nil },
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			chatService := tt.chatServiceMock(mc)
			messageService := tt.messageServiceMock(mc)
			api := chat.NewImplementation(chatService, messageService)

			resp, err := api.CreateChat(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})

	}
}
