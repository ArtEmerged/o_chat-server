package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_chat-server/internal/service/chat"
)

func TestDeleteChat(t *testing.T) {
	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepo
	type txManagerMock func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		mc            = minimock.NewController(t)
		ctx           = context.Background()
		chatID        = gofakeit.Int64()
		repositoryErr = errors.New("repository error")
	)

	tests := []struct {
		name          string
		args          args
		want          int64
		wantErr       error
		chatRepoMock  chatRepoMockFunc
		txManagerMock txManagerMock
	}{
		{
			name: "success create chat",
			args: args{
				ctx: ctx,
				req: chatID,
			},
			wantErr: nil,

			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepo {
				repoMock := mocks.NewChatRepoMock(mc)
				repoMock.DeleteChatMock.Expect(ctx, chatID).Return(nil)
				return repoMock
			},

			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: chatID,
			},
			wantErr: repositoryErr,

			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepo {
				repoMock := mocks.NewChatRepoMock(mc)
				repoMock.DeleteChatMock.Expect(ctx, chatID).Return(repositoryErr)
				return repoMock
			},

			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			chatRepo := tt.chatRepoMock(mc)
			txManager := tt.txManagerMock(mc)
			service := chat.New(chatRepo, txManager)
			err := service.DeleteChat(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.wantErr, err)
		})
	}
}
