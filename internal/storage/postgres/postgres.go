package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const fn = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXIST tasks(
	    id INTEGER PRIMARY KEY,
	    task_name TEXT NOT NULL,
	    description TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_task_name ON tasks(task_name);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Storage{db: db}, nil
}
