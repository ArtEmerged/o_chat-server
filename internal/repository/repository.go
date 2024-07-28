package repository

import (
	"github.com/ArtEmerged/o_chat-server/internal/definitions"
	"github.com/jackc/pgx/v4/pgxpool"
)

type chatRepo struct {
	db *pgxpool.Pool
}

var _ definitions.ChatRepo = (*chatRepo)(nil)

func New(db *pgxpool.Pool) *chatRepo {
	return &chatRepo{db: db}
}
