package datastore

import (
	"context"
	"github.com/danyazab/animal/internal/animal/model"
	"github.com/danyazab/animal/pkg/database/core"
)

type catRepository struct {
	db core.NamedExecutor
}

func NewCatRepository(db core.NamedExecutor) model.CatRepository {
	return &catRepository{
		db: db,
	}
}

func (r catRepository) Update(ctx context.Context, entity model.Cat) error {
	query :=
		`UPDATE cat SET
			name = :name,
			description = :description,
			breed = :breed,
			birthday = :birthday,
			sex = :sex,
			tail_length = :tail_length,
			color = :color,
			wool_type = :wool_type,
			is_chipped = :is_chipped,
			weight = :weight,
			updated_at = NOW()
		WHERE id = :id`

	return core.Execute(ctx, r.db, query, entity)
}

func (r catRepository) Store(ctx context.Context, entity model.Cat) (model.Cat, error) {
	res, _, err := core.Single[model.Cat](
		ctx,
		r.db,
		`INSERT INTO cat (name, description, breed, birthday, sex, tail_length, color, wool_type, is_chipped, weight, created_at, updated_at)
				VALUES (:name, :description, :breed, :birthday, :sex, :tail_length, :color, :wool_type, :is_chipped, :weight, NOW(), NOW()) RETURNING *`,
		entity,
	)

	return res, err
}

func (r catRepository) FindByID(ctx context.Context, id uint) (model.Cat, bool, error) {
	query := "SELECT * FROM cat WHERE id = :id"
	arg := map[string]interface{}{
		"id": id,
	}

	return core.Single[model.Cat](ctx, r.db, query, arg)
}
