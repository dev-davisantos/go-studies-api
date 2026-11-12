package database

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func initDatabase() error {
	dsb := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("pgx", dsb)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return err
	}

	var exists bool
	err = db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)",
		"go_studies_api_db",
	).Scan(&exists)

	if !exists {
		if _, err := db.Exec("CREATE DATABASE go_category_api_db"); err != nil {
			return err
		}
	}
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func NewPostgresDB() (*sql.DB, error) {
	if initErr := initDatabase(); initErr != nil {
		return nil, initErr
	}
	dsb := "postgres://postgres:postgres@localhost:5432/go_category_api_db?sslmode=disable"

	db, err := sql.Open("pgx", dsb)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	schema, error := os.ReadFile("script/schema.sql")
	if error != nil {
		return nil, error
	}

	if _, err := db.Exec(string(schema)); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
