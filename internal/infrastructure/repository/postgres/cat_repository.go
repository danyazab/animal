package postgres

import (
	"context"
	"danyazab/animal/internal/animal/model"
	"danyazab/animal/pkg/database/core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type catRepository struct {
	db *sqlx.DB
}

func NewCatRepository(db *sqlx.DB) model.CatRepository {
	return &catRepository{
		db: db,
	}
}

func (r catRepository) Store(ctx context.Context, entity model.Cat) (model.Cat, error) {
	res, _, err := core.Single[model.Cat](
		ctx,
		r.db,
		fmt.Sprintf(
			`INSERT INTO cat (name, description, breed, birthday, sex, tail_length, color, wool_type, is_chipped, weight, created_at, updated_at)
				VALUES (:name, :description, :breed, :birthday, :sex, :tail_length, :color, :wool_type, :is_chipped, :weight, NOW(), NOW()) RETURNING *`,
		), entity)

	return res, err
}

func (r catRepository) GetAll(ctx context.Context) ([]model.Cat, uint64, error) {
	return core.Paged[model.Cat](
		ctx,
		r.db,
		"SELECT * FROM cat",
		"SELECT COUNT(*) FROM cat",
	)
}
