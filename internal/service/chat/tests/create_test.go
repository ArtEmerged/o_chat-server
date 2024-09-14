package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	txMocks "github.com/ArtEmerged/o_chat-server/internal/client/db/mocks"
	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_chat-server/internal/service/chat"
)


// TODO fix tx manager 
func TestCreateChat(t *testing.T) {
	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepo
	type txManagerMock func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.CreateChatRequest
	}

	var (
		mc       = minimock.NewController(t)
		ctx      = context.Background()
		chatID   = int64(777)
		chatName = gofakeit.StreetName()
		userIDs  = []int64{
			1,
			4,
			3,
			2,
		}
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
				req: &model.CreateChatRequest{
					ChatName:  chatName,
					CreatorID: 5,
					UserIDs:   userIDs,
				},
			},
			want:    chatID,
			wantErr: nil,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepo {
				repoMock := mocks.NewChatRepoMock(mc)
				req := &model.CreateChatRequest{
					ChatName:  chatName,
					CreatorID: 5,
					UserIDs:   []int64{1, 2, 3, 4, 5},
				}
				repoMock.CreateChatMock.Expect(ctx, req).Return(chatID, nil)
				repoMock.AddUsersToChatMock.Expect(ctx, chatID, req.UserIDs).Return(nil)
				return repoMock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				txManagerMock := txMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommittedMock.Expect(ctx, func(ctx context.Context) error { return nil }).Return(nil)
				return txManagerMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			chatRepo := tt.chatRepoMock(mc)
			txManager := tt.txManagerMock(mc)

			service := chat.New(chatRepo, txManager)
			got, err := service.CreateChat(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err)
		})
	}
}
