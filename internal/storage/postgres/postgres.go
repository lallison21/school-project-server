package postgres

import (
	_ "github.com/lib/pq"

	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lallison21/to-do/internal/config"
)

type Storage struct {
	db *sqlx.DB
}

func New(storageConfig config.StorageConfig) (*Storage, error) {
	const fn = "storage.postgres.New"

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s post=%s user=%s dbname=%s password=%s sslmode=%s",
		storageConfig.Host, storageConfig.Port, storageConfig.Username, storageConfig.DBName, storageConfig.Password, storageConfig.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return nil, nil
}
