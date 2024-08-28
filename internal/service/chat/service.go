package chat

import (
	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/service"
)

type chatService struct {
	repo      repository.ChatRepo
	txManager db.TxManager
}

// New returns new chat service
func New(repo repository.ChatRepo, txManager db.TxManager) service.ChatService {
	return &chatService{
		repo:      repo,
		txManager: txManager}
}
