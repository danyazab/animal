package datastore

import (
	"danyazab/animal/pkg/database/dbtesting"
	"danyazab/animal/pkg/database/migrator"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func FixturesDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../fixtures")
}

func TestMain(m *testing.M) {
	dbname := "testdb"
	migrationsDir := "file://../migrations"

	config := dbtesting.Config{
		ExposePort: "5432/tcp",
		DBName:     "testdb",
		Image:      "postgres:14.3-alpine",
		Env: []string{
			fmt.Sprintf("POSTGRES_DB=%s", dbname),
			"POSTGRES_PASSWORD=password",
			"POSTGRES_USER=postgres",
		},
	}

	os.Exit(dbtesting.RunContainer(
		m,
		config,
		func(conn *sql.DB, dbname string) error {
			return migrator.MigratePostgres(conn, dbname, migrationsDir, migrator.UpCmd)
		},
	))
}
