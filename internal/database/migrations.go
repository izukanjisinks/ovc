package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrations(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose dialect: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	log.Println("migrations applied successfully")
	return nil
}

func RollbackMigration(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)
	goose.SetDialect("postgres")
	return goose.Down(db, "migrations")
}

func ResetDatabase(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)
	goose.SetDialect("postgres")
	return goose.Reset(db, "migrations")
}

func MigrationStatus(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)
	goose.SetDialect("postgres")
	return goose.Status(db, "migrations")
}
