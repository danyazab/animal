package internal

import (
	"danyazab/animal/config"
	"danyazab/animal/internal/infrastructure/repository/postgres"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
)

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
	if err := c.Provide(sqlxProvider); err != nil {
		return nil, err
	}
	if err := c.Provide(postgres.NewCatRepository); err != nil {
		return nil, err
	}

	return c, nil
}

func sqlxProvider(cfg *config.Database) *sqlx.DB {
	connUrl := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name)

	db := sqlx.MustConnect("pgx", connUrl)

	return db
}
