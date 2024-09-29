package message

import (
	"github.com/ArtEmerged/library/client/cache"
	dbClient "github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_chat-server/internal/repository"
)

type messageRepo struct {
	db    dbClient.Client
	cache cache.Cache
}

// New creates a new instance of messageRepo with the given database connection pool.
// db - database connection pool
func New(db dbClient.Client, cache cache.Cache) repository.MessageRepo {
	return &messageRepo{db: db, cache: cache}
}
