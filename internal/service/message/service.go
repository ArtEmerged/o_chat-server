package message

import (
	"github.com/ArtEmerged/library/client/cache"
	"github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/service"
)

type messageService struct {
	repo      repository.MessageRepo
	cache     cache.Cache
	txManager db.TxManager
}

// New returns new chat service
func New(repo repository.MessageRepo, cache cache.Cache, txManager db.TxManager) service.MessageService {

	return &messageService{
		repo:      repo,
		cache:     cache,
		txManager: txManager,
	}
}
