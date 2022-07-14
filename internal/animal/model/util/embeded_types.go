package util

import (
	"time"
)

type (
	TimeStamps struct {
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)
