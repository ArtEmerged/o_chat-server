package service

import (
	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
)

type chatService struct {
	repo def.ChatRepo
}

var _ def.ChatService = (*chatService)(nil)

func New(repo def.ChatRepo) *chatService {
	return &chatService{repo: repo}
}
