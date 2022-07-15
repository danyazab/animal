package core

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// sqlxConnection may be unnecessary
type sqlxConnection interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	ExecuteContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type sqlxWrapper struct {
	connection NamedExecutor
}

func (a sqlxWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	var arg interface{}
	if len(args) > 0 {
		arg = args[0]
	} else {
		arg = map[string]interface{}{}
	}

	statement, err := a.connection.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil
	}

	return statement.QueryRowxContext(ctx, arg)
}

func (a sqlxWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	var arg interface{}
	if len(args) > 0 {
		arg = args[0]
	} else {
		arg = map[string]interface{}{}
	}

	statement, err := a.connection.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return statement.QueryxContext(ctx, arg)
}

func (a sqlxWrapper) ExecuteContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var arg interface{}
	if len(args) > 0 {
		arg = args[0]
	} else {
		arg = map[string]interface{}{}
	}

	statement, err := a.connection.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return statement.ExecContext(ctx, arg)
}
