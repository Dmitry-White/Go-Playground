package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const MIGRATIONS_URL = "file://cmd/migrations"
const DSN = "postgresql://postgres:password@localhost:5432/users?sslmode=disable&timezone=UTC&connect_timeout=5"

func main() {
	m, err := migrate.New(MIGRATIONS_URL, DSN)
	if err != nil {
		log.Fatalln(err)
	}

	initialVersion, _, _ := m.Version()
	log.Println("Initial version:", initialVersion)

	migrationErr := m.Up()
	if migrationErr != nil {
		log.Fatalln(migrationErr)
	}

	currentVersion, _, _ := m.Version()
	log.Println("Current version:", currentVersion)
}
