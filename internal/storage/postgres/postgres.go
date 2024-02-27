package postgres

import (
	_ "github.com/lib/pq"

	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lallison21/school-project-server/internal/config"
)

type Storage struct {
	db *sqlx.DB
}

func New(storageConfig config.StorageConfig) (*Storage, error) {
	const fn = "storage.postgres.New"

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		storageConfig.Host, storageConfig.Port, storageConfig.Username, storageConfig.DBName, storageConfig.Password, storageConfig.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateRole(roleName string, accessLevel int) (int64, error) {
	const fn = "storage.postgres.CreateRole"

	stmt, err := s.db.Prepare("INSERT INTO role_list(role_name, access_level) VALUES($1, $2) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", fn, err)
	}

	var id int64
	err = stmt.QueryRow(roleName, accessLevel).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last index id: %w", fn, err)
	}

	return id, nil
}

func (s *Storage) GetRoleById(id int) (string, error) {
	const fn = "storage.postgres.GetRoleById"

	stmt, err := s.db.Prepare("SELECT role_name FROM role_list WHERE id = $1")
	if err != nil {
		return "", fmt.Errorf("%s: prepare statement: %w", fn, err)
	}

	var roleName string
	if err = stmt.QueryRow(id).Scan(&roleName); err != nil {
		return "", fmt.Errorf("%s: execute statement: %w", fn, err)
	}

	return roleName, nil
}

// TODO: implement method:
// func (s *Storage) DeleteRole(id int) error
