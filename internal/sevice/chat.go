package service

import (
	"context"

	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
)

func (s *chatService) CreateChat(ctx context.Context, in *def.CreateChatRequest) (id int64, err error) {
	in.UserIDs = append(in.UserIDs, in.CreatorID)

	in.UserIDs = uniqueSliceInt64(in.UserIDs)

	return s.repo.CreateChat(ctx, in)
}
func (s *chatService) DeleteChat(ctx context.Context, id int64) error {
	return s.repo.DeleteChat(ctx, id)
}
func (s *chatService) SendMessage(ctx context.Context, in *def.SendMessageRequest) error {
	return s.repo.SendMessage(ctx, in)
}
