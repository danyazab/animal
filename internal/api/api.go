package api

import (
	"danyazab/animal/config"
	"danyazab/animal/internal/api/controller/dog"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"go.uber.org/dig"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationsDirPath = "file://database/migration"

type API struct {
	dig.In

	dog.Controller
}

func RunServer(api API, db *sqlx.DB, cfg *config.Database) error {
	if err := runMigrations(db, cfg.Name); err != nil {
		return err
	}
	port := 8000

	app := echo.New()
	api.initRoutes(app)

	return app.Start(fmt.Sprintf(":%d", port))
}

func runMigrations(db *sqlx.DB, dbName string) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsDirPath, dbName, driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (a *API) initRoutes(app *echo.Echo) {
	app.POST("/pet/dog", a.Create)
	app.GET("/pet/dog", a.List)
}
