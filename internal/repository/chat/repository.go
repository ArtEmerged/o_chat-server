package chat

import (
	dbClient "github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
)

type chatRepo struct {
	db dbClient.Client
}

// New creates a new instance of chatRepo with the given database connection pool.
// db - database connection pool
func New(db dbClient.Client) repository.ChatRepo {
	return &chatRepo{db: db}
}
