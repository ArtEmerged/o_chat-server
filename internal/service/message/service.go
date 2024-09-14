package message

import (
	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/service"
)

type messageService struct {
	repo      repository.MessageRepo
	txManager db.TxManager
}

// New returns new chat service
func New(repo repository.MessageRepo, txManager db.TxManager) service.MessageService {
	return &messageService{
		repo:      repo,
		txManager: txManager,
	}
}
