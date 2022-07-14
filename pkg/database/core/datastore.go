package core

import (
	"context"
	"github.com/jmoiron/sqlx"
)

// NamedExecutor interface for Executor operations
type NamedExecutor interface {
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
}
