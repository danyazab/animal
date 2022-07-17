package model

import (
	"context"
	util2 "danyazab/animal/internal/animal/model/util"
	"time"
)

//go:generate  mockery --dir . --name CatRepository --output ./../../gen/mock
type CatRepository interface {
	Store(ctx context.Context, entity Cat) (Cat, error)
	GetAll(ctx context.Context, limit uint) ([]Cat, uint64, error)
}

type Cat struct {
	ID          uint           `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	Breed       string         `db:"breed"`
	Birthday    time.Time      `db:"birthday"`
	Sex         util2.TypeSex  `db:"sex"`
	TailLength  uint           `db:"tail_length"`
	Color       string         `db:"color"`
	WoolType    util2.TypeWool `db:"wool_type"`
	IsChipped   bool           `db:"is_chipped"`
	Weight      float32        `db:"weight"`
	util2.TimeStamps
}
