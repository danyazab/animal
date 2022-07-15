package core

import (
	"context"
	"fmt"
	"strings"
)

func wrapConnection(executor NamedExecutor) sqlxConnection {
	return sqlxWrapper{
		connection: executor,
	}
}

func Execute(ctx context.Context, db NamedExecutor, query string, args ...interface{}) error {
	_, err := wrapConnection(db).ExecuteContext(ctx, query, args...)

	return err
}

func AdjustQueryWhereInStatement[TEntity any](query string, columnName string, args []TEntity) string {
	if len(args) == 0 {
		return query
	}

	values := make([]string, 0, len(args))
	for i := range args {
		values = append(values, fmt.Sprintf(`:%s_%d`, columnName, i))
	}

	statement := strings.Join(values, ", ")

	return strings.Replace(query, `:`+columnName, statement, 1)
}

func ExtendArgsWithInParams[TEntity any](args map[string]interface{}, columnName string, values []TEntity) map[string]interface{} {
	for i := range values {
		args[fmt.Sprintf(`%s_%d`, columnName, i)] = values[i]
	}

	return args
}

func Count(ctx context.Context, db NamedExecutor, query string, args ...interface{}) (count uint64, err error) {
	err = wrapConnection(db).QueryRowContext(ctx, query, args...).Scan(&count)

	return
}

func Single[TEntity any](ctx context.Context, db NamedExecutor, query string, args ...interface{}) (result TEntity, found bool, err error) {
	items, err := Multiple[TEntity](ctx, db, query, args...)
	if err != nil {
		return
	}

	if len(items) == 0 {
		found = false

		return
	}

	found = true

	if len(items) > 1 {
		err = fmt.Errorf(`multiple rows were returned`)

		return
	}

	result = items[0]

	return
}

func Multiple[TEntity any](ctx context.Context, db NamedExecutor, query string, args ...interface{}) ([]TEntity, error) {
	rows, err := wrapConnection(db).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []TEntity

	for rows.Next() {
		var record TEntity
		if err = rows.StructScan(&record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func MultipleScalar[TEntity any](ctx context.Context, db NamedExecutor, query string, args ...interface{}) ([]TEntity, error) {
	rows, err := wrapConnection(db).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []TEntity

	for rows.Next() {
		var record TEntity
		if err = rows.Scan(&record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func Insert(ctx context.Context, db NamedExecutor, query string, args ...interface{}) error {
	rows, err := wrapConnection(db).QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return rows.Close()
}

func Remove(ctx context.Context, db NamedExecutor, query string, args ...interface{}) error {
	rows, err := wrapConnection(db).QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	if rows != nil {
		err = rows.Close()
	}

	return err
}

func Paged[TEntity any](ctx context.Context, db NamedExecutor, query, countQuery string, args ...interface{}) ([]TEntity, uint64, error) {
	records, err := Multiple[TEntity](ctx, db, query, args...)
	if err != nil {
		return nil, 0, err
	}

	count, err := Count(ctx, db, countQuery, args...)

	return records, count, err
}
