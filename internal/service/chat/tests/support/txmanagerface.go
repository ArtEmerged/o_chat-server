package support

import (
	"context"

	"github.com/ArtEmerged/library/client/db"
)

type txManagerFake struct {
}

// NewTxManagerFake fake tx manager
func NewTxManagerFake() db.TxManager {
	return &txManagerFake{}
}

func (tx *txManagerFake) ReadCommitted(ctx context.Context, f db.Handler) error {
	return f(ctx)
}
