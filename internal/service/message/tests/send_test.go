package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_chat-server/internal/service/message"
)

func TestSendMessage(t *testing.T) {
	type messageRepoMockFunc func(mc *minimock.Controller) repository.MessageRepo
	type txManagerMock func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.SendMessageRequest
	}

	var (
		mc            = minimock.NewController(t)
		ctx           = context.Background()
		chatID        = gofakeit.Int64()
		fromUserID    = gofakeit.Int64()
		text          = gofakeit.Sentence(30)
		repositoryErr = errors.New("repository error")
	)

	tests := []struct {
		name            string
		args            args
		want            int64
		wantErr         error
		messageRepoMock messageRepoMockFunc
		txManagerMock   txManagerMock
	}{
		{
			name: "success create chat",
			args: args{
				ctx: ctx,
				req: &model.SendMessageRequest{
					ChatID:     chatID,
					FromUserID: fromUserID,
					Text:       text,
				},
			},
			wantErr: nil,
			messageRepoMock: func(mc *minimock.Controller) repository.MessageRepo {
				repoMock := mocks.NewMessageRepoMock(mc)
				req := &model.SendMessageRequest{
					ChatID:     chatID,
					FromUserID: fromUserID,
					Text:       text,
				}
				repoMock.SendMessageMock.Expect(ctx, req).Return(nil)
				return repoMock
			},

			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "validate error",
			args: args{
				ctx: ctx,
				req: &model.SendMessageRequest{
					ChatID:     chatID,
					FromUserID: fromUserID,
				},
			},
			wantErr:         fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field text is required"),
			messageRepoMock: func(mc *minimock.Controller) repository.MessageRepo { return nil },
			txManagerMock:   func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: &model.SendMessageRequest{
					ChatID:     chatID,
					FromUserID: fromUserID,
					Text:       text,
				},
			},
			wantErr: repositoryErr,

			messageRepoMock: func(mc *minimock.Controller) repository.MessageRepo {
				repoMock := mocks.NewMessageRepoMock(mc)
				req := &model.SendMessageRequest{
					ChatID:     chatID,
					FromUserID: fromUserID,
					Text:       text,
				}
				repoMock.SendMessageMock.Expect(ctx, req).Return(repositoryErr)
				return repoMock
			},

			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			messageRepo := tt.messageRepoMock(mc)
			txManager := tt.txManagerMock(mc)

			service := message.New(messageRepo, txManager)
			err := service.SendMessage(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.wantErr, err)
		})
	}
}
