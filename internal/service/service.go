package service

import (
	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
)

type chatService struct {
	repo def.ChatRepo
}

// New returns new chat service
func New(repo def.ChatRepo) def.ChatService {
	return &chatService{repo: repo}
}
