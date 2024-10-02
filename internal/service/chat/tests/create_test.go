package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/ArtEmerged/library/client/cache"
	cacheMock "github.com/ArtEmerged/library/client/cache/mocks"
	"github.com/ArtEmerged/library/client/db"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_chat-server/internal/service/chat"
	testSupport "github.com/ArtEmerged/o_chat-server/internal/service/chat/tests/support"
)

func TestCreateChat(t *testing.T) {
	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepo
	type cacheMockFunc func(mc *minimock.Controller) cache.Cache

	type args struct {
		ctx context.Context
		req *model.CreateChatRequest
	}

	var (
		mc            = minimock.NewController(t)
		ctx           = context.Background()
		chatID        = int64(777)
		chatName      = gofakeit.StreetName()
		repositoryErr = errors.New("repository error")
		userIDs       = []int64{
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
		cacheMock     cacheMockFunc
		txManagerFace db.TxManager
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
			cacheMock: func(mc *minimock.Controller) cache.Cache {
				mock := cacheMock.NewCacheMock(mc)

				newChat := &model.Chat{
					ID:        chatID,
					ChatName:  chatName,
					CreatorID: 5,
					UserIDs:   []int64{1, 2, 3, 4, 5},
				}
				mock.SetMock.Expect(ctx, model.ChatCacheKey(chatID), newChat, 0).Return(nil)

				return mock
			},
			txManagerFace: testSupport.NewTxManagerFake(),
		},
		{
			name: "repository error create chat",
			args: args{
				ctx: ctx,
				req: &model.CreateChatRequest{
					ChatName:  chatName,
					CreatorID: 5,
					UserIDs:   userIDs,
				},
			},
			want:    -1,
			wantErr: repositoryErr,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepo {
				repoMock := mocks.NewChatRepoMock(mc)
				req := &model.CreateChatRequest{
					ChatName:  chatName,
					CreatorID: 5,
					UserIDs:   []int64{1, 2, 3, 4, 5},
				}
				repoMock.CreateChatMock.Expect(ctx, req).Return(-1, repositoryErr)
				return repoMock
			},
			cacheMock:     func(mc *minimock.Controller) cache.Cache { return nil },
			txManagerFace: testSupport.NewTxManagerFake(),
		},
		{
			name: "repository error add users to chat",
			args: args{
				ctx: ctx,
				req: &model.CreateChatRequest{
					ChatName:  chatName,
					CreatorID: 5,
					UserIDs:   userIDs,
				},
			},
			want:    -1,
			wantErr: repositoryErr,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepo {
				repoMock := mocks.NewChatRepoMock(mc)
				req := &model.CreateChatRequest{
					ChatName:  chatName,
					CreatorID: 5,
					UserIDs:   []int64{1, 2, 3, 4, 5},
				}
				repoMock.CreateChatMock.Expect(ctx, req).Return(chatID, nil)
				repoMock.AddUsersToChatMock.Expect(ctx, chatID, req.UserIDs).Return(repositoryErr)
				return repoMock
			},
			cacheMock:     func(mc *minimock.Controller) cache.Cache { return nil },
			txManagerFace: testSupport.NewTxManagerFake(),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			chatRepo := tt.chatRepoMock(mc)
			txManager := tt.txManagerFace
			cache := tt.cacheMock(mc)

			service := chat.New(chatRepo, txManager, cache)
			got, err := service.CreateChat(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err)
		})
	}
}
