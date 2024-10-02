package chat

import (
	"github.com/ArtEmerged/library/client/cache"
	"github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_chat-server/internal/repository"
	"github.com/ArtEmerged/o_chat-server/internal/service"
)

type chatService struct {
	repo      repository.ChatRepo
	txManager db.TxManager
	cache     cache.Cache
}

// New returns new chat service
func New(repo repository.ChatRepo, txManager db.TxManager, cache cache.Cache) service.ChatService {
	return &chatService{
		repo:      repo,
		txManager: txManager,
		cache:     cache,
	}
}
