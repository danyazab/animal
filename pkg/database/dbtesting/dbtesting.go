package dbtesting

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-txdb"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
)

// Migrator performs database migrations
type Migrator func(conn *sql.DB, dbname string) error

// Config for a db container
type Config struct {
	ExposePort string
	DBName     string
	Image      string
	Env        []string
}

// RunContainer spins up a docker container for integration testing
func RunContainer(m *testing.M, config Config, migrator Migrator) int {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	image := strings.Split(config.Image, ":")
	if len(image) != 2 {
		log.Fatalf("invalid image string: %s", config.Image)
	}

	resource, err := pool.Run(image[0], image[1], config.Env)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	var db *sql.DB

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	var postgresDSN string
	if err = pool.Retry(func() error {
		postgresDSN = fmt.Sprintf(
			"postgres://postgres:password@localhost:%s/%s?sslmode=disable",
			resource.GetPort(config.ExposePort),
			config.DBName,
		)
		db, err = sql.Open("pgx", postgresDSN)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	err = migrator(db, config.DBName)
	if err != nil {
		log.Fatal(err)
	}

	txdb.Register("txdb", "pgx", postgresDSN)

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	return code
}

// Inject provides a transaction within test container
func Inject(f func(*testing.T, *sqlx.DB)) func(*testing.T) {
	return func(t *testing.T) {
		db, err := sql.Open("txdb", "identifier")
		if err != nil {
			panic(err)
		}
		defer func() {
			err := db.Close()
			if err != nil {
				t.Logf("rollback failed during normal execution: %v", err)
			}
		}()

		xDB := sqlx.NewDb(db, "postgres")

		// mark this function as a helper function for stack introspection
		t.Helper()

		f(t, xDB)
	}
}
