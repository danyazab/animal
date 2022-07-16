package internal

import (
	"danyazab/animal/config"
	"danyazab/animal/internal/infrastructure/repository/postgres/datastore"
	"danyazab/animal/pkg/database/core"
	"danyazab/animal/pkg/database/migrator"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
)

const migrationsDirPath = "file://internal/infrastructure/repository/postgres/migrations"

func Invoke(fn, cfg interface{}) error {
	c, err := container(cfg)
	if err != nil {
		return err
	}
	return c.Invoke(fn)
}

func container(cfg interface{}) (*dig.Container, error) {
	c := dig.New()
	if err := c.Provide(func() *dig.Container { return c }); err != nil {
		return nil, err
	}
	if err := c.Provide(cfg); err != nil {
		return nil, err
	}
	if err := c.Provide(SqlxProvider); err != nil {
		return nil, err
	}
	if err := c.Provide(datastore.NewCatRepository); err != nil {
		return nil, err
	}

	return c, nil
}

func SqlxProvider(cfg *config.Database) core.NamedExecutor {
	connUrl := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name)

	db := sqlx.MustConnect("pgx", connUrl)

	if err := migrator.MigratePostgres(db.DB, cfg.Name, migrationsDirPath, migrator.UpCmd); err != nil {
		panic(err)
	}

	return db
}
