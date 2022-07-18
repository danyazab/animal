package internal

import (
	"danyazab/animal/config"
	"danyazab/animal/internal/infrastructure/network/petfinder"
	"danyazab/animal/internal/infrastructure/repository/postgres/datastore"
	"danyazab/animal/pkg/database/core"
	"danyazab/animal/pkg/database/migrator"
	"danyazab/animal/pkg/http/client"
	"fmt"
	"github.com/go-resty/resty/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
	"time"
)

const migrationsDirPath = "file://internal/infrastructure/repository/postgres/migrations"

func Invoke(fn, cfg interface{}) error {
	c, err := container(cfg)
	if err != nil {
		return err
	}
	return c.Invoke(fn)
}

func providers() []interface{} {
	return []interface{}{
		sqlxProvider,
		restyClientProvider,
		datastore.NewCatRepository,
		petfinder.NewClient,
		client.NewTransport,
	}
}

func container(cfg interface{}) (*dig.Container, error) {
	c := dig.New()
	if err := c.Provide(func() *dig.Container { return c }); err != nil {
		return nil, err
	}
	if err := c.Provide(cfg); err != nil {
		return nil, err
	}

	for _, provider := range providers() {
		if err := c.Provide(provider); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func sqlxProvider(cfg *config.Database) core.NamedExecutor {
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

func restyClientProvider() *resty.Client {
	return resty.New().
		SetTimeout(10 * time.Second).
		SetBaseURL("https://api.petfinder.com")
}
