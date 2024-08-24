package chat

import (
	dbClient "github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
)

type messageRepo struct {
	db dbClient.Client
}

// New creates a new instance of messageRepo with the given database connection pool.
// db - database connection pool
func New(db dbClient.Client) repository.MessageRepo {
	return &messageRepo{db: db}
}
