package datastore

import (
	"context"
	"danyazab/animal/internal/animal/model"
	"danyazab/animal/internal/animal/model/util"
	"danyazab/animal/pkg/database/dbtesting"
	"danyazab/animal/pkg/database/fixtures"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var newCat = model.Cat{
	Name:       "test",
	Birthday:   time.Time{},
	Sex:        util.TypeSexFemale,
	TailLength: 10,
	Color:      "red",
	WoolType:   util.TypeWoolLong,
	IsChipped:  false,
	Weight:     4.5,
	TimeStamps: util.TimeStamps{},
}

func Test_Store(t *testing.T) {
	t.Run("successfully stored new cat into DB", dbtesting.Inject(func(t *testing.T, db *sqlx.DB) {
		exampleRepo := NewCatRepository(db)

		cat, err := exampleRepo.Store(context.Background(), newCat)
		assert.NoError(t, err)
		assert.Greater(t, cat.ID, uint(0))
	}))
}

func Test_FindByID(t *testing.T) {
	type testCase = struct {
		name       string
		catID      uint
		exist      bool
		tablesDump []string
	}

	fDir := FixturesDir()
	cases := []testCase{
		{
			name:  "Cat with id:3 exist in DB",
			catID: 3,
			exist: true,
			tablesDump: []string{
				fDir + "/cat_id_3.yaml",
			},
		},
		{
			name:  "Cat with id:3 not exist in DB",
			catID: 3,
			exist: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d case: %s", i, c.name), dbtesting.Inject(func(t *testing.T, db *sqlx.DB) {
			fixtures.PrepareTestDatabase(db, c.tablesDump...)
			exampleRepo := NewCatRepository(db)

			cat, found, err := exampleRepo.FindByID(context.Background(), c.catID)
			assert.NoError(t, err)
			assert.Equal(t, c.exist, found)
			if c.exist {
				assert.Equal(t, cat.ID, c.catID)
			}
		}))
	}
}

func Test_Update(t *testing.T) {
	type testCase = struct {
		name       string
		cat        model.Cat
		compare    func(c model.Cat, f model.CatRepository)
		tablesDump []string
	}

	fDir := FixturesDir()
	cases := []testCase{
		{
			name: "update cat id:3 successfully",
			cat: model.Cat{
				ID:       3,
				Name:     "Test",
				Sex:      util.TypeSexFemale,
				WoolType: util.TypeWoolShort,
			},
			tablesDump: []string{
				fDir + "/cat_id_3.yaml",
			},
			compare: func(c model.Cat, f model.CatRepository) {
				cat, _, _ := f.FindByID(context.Background(), c.ID)
				// skip date time fields
				cat.UpdatedAt = c.UpdatedAt
				cat.CreatedAt = c.CreatedAt

				cj, _ := json.Marshal(c)
				cj1, _ := json.Marshal(cat)
				assert.JSONEq(t, string(cj), string(cj1))
			},
		},
		{
			name: "update cat not existed cat, id:31 successfully",
			cat: model.Cat{
				ID:       31,
				Name:     "Test",
				Sex:      util.TypeSexFemale,
				WoolType: util.TypeWoolShort,
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d case: %s", i, c.name), dbtesting.Inject(func(t *testing.T, db *sqlx.DB) {
			fixtures.PrepareTestDatabase(db, c.tablesDump...)
			exampleRepo := NewCatRepository(db)

			err := exampleRepo.Update(context.Background(), c.cat)
			assert.NoError(t, err)

			if c.compare != nil {
				c.compare(c.cat, exampleRepo)
			}
		}))
	}
}
