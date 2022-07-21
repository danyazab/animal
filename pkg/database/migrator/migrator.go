package migrator

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // driver for working with postgresql
)

// CmdFunc represents specific migration command
type CmdFunc func(m *migrate.Migrate) error

// MigratePostgres migrates postgres database with a given command
func MigratePostgres(db *sql.DB, dbName, path string, cmd CmdFunc) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: dbName,
	})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	return cmd(m)
}

// UpCmd performs migrations up to the most recent version
func UpCmd(m *migrate.Migrate) error {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrations failed: %w", err)
	}

	return nil
}

// DownCmd performs migrations down
func DownCmd(m *migrate.Migrate) error {
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrations failed: %w", err)
	}

	return nil
}

// ToVersionCmd looks at the currently active migration version,
// then migrates either up or down to the specified version.
func ToVersionCmd(version uint) CmdFunc {
	return func(m *migrate.Migrate) error {
		if err := m.Migrate(version); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("migrations failed: %w", err)
		}

		return nil
	}
}
