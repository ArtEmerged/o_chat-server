package chat

import (
	"github.com/ArtEmerged/library/client/cache"
	dbClient "github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_chat-server/internal/repository"
)

type chatRepo struct {
	db    dbClient.Client
	cache cache.Cache
}

// New creates a new instance of chatRepo with the given database connection pool.
// db - database connection pool
func New(db dbClient.Client, cache cache.Cache) repository.ChatRepo {
	return &chatRepo{db: db, cache: cache}
}
