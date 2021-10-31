package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB) {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://repository/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	version, _, _ := m.Version()
	fmt.Println("version: ", version)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		version2, _, _ := m.Version()
		fmt.Println("version_Second: ", version2)

		log.Print(err)

		gap := version - version2
		m.Steps(int(gap))

		m.Force(int(version))
	}
}
