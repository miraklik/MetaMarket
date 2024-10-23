package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	// ! Библиотека для Миграций
	"github.com/golang-migrate/migrate/v4"
	// ! Миграция с SQLite
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	// ! Миграция с Файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}

		if dirtyErr, ok := err.(*migrate.ErrDirty); ok {
			fmt.Printf("Database is in a dirty state (version: %d). Forcing migration version...\n", dirtyErr.Version)
			if forceErr := m.Force(dirtyErr.Version); forceErr != nil {
				log.Fatalf("failed to force migration version: %v", forceErr)
			}
			fmt.Println("Dirty state fixed, attempting migration again...")

			if retryErr := m.Up(); retryErr != nil {
				log.Fatalf("migration failed on retry: %v", retryErr)
			}
			fmt.Println("Migrations applied successfully after fixing dirty state.")
			return
		}

		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("migrations applied successfully")
}
