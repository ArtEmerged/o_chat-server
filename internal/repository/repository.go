package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"

	def "github.com/ArtEmerged/o_chat-server/internal/definitions"
)

type chatRepo struct {
	db *pgxpool.Pool
}

// New creates a new instance of chatRepo with the given database connection pool.
// db - pointer to the PostgreSQL connection pool
func New(db *pgxpool.Pool) def.ChatRepo {
	return &chatRepo{db: db}
}
