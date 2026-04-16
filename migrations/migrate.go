package migrations

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations chạy migration khi app start
func RunMigrations(dbURL string) {
	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)

	if err != nil {
		log.Fatal("migration init error:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("migration up error:", err)
	}

	log.Println("Database migrated successfully")
}
