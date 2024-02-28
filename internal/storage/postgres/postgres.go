package postgres

import (
	"github.com/lallison21/school-project-server/internal/http-server/handlers/url/role"
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

func (s *Storage) GetRoles() ([]role.Role, error) {
	const fn = "storage.postgres.GetRoleById"
	var res []role.Role

	stmt, err := s.db.Prepare("SELECT * FROM role_list")
	if err != nil {
		return res, fmt.Errorf("%s: prepare statement: %w", fn, err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return res, fmt.Errorf("%s: execute statement: %w", fn, err)
	}

	for rows.Next() {
		var r role.Role
		if err := rows.Scan(&r.Id, &r.RoleName, &r.AccessLeve); err != nil {
			return res, fmt.Errorf("%s: scan error: %w", fn, err)
		}
		res = append(res, r)
	}
	if err := rows.Err(); err != nil {
		return res, fmt.Errorf("%s: rows error: %w", fn, err)
	}
	if err := rows.Close(); err != nil {
		return res, fmt.Errorf("%s: close rows: %w", fn, err)
	}

	return res, nil
}

func (s *Storage) GetRoleById(id int) (role.Role, error) {
	const fn = "storage.postgres.GetRoleById"
	var res role.Role

	stmt, err := s.db.Prepare("SELECT * FROM role_list WHERE id = $1")
	if err != nil {
		return res, fmt.Errorf("%s: prepare statement: %w", fn, err)
	}

	if err = stmt.QueryRow(id).Scan(&res.Id, &res.RoleName, &res.AccessLeve); err != nil {
		return res, fmt.Errorf("%s: execute statement: %w", fn, err)
	}

	return res, nil
}

func (s *Storage) CreateRole(roleName string, accessLevel int) (role.Role, error) {
	const fn = "storage.postgres.CreateRole"
	var res role.Role

	stmt, err := s.db.Prepare("INSERT INTO role_list(role_name, access_level) VALUES($1, $2) RETURNING *")
	if err != nil {
		return res, fmt.Errorf("%s: prepare statement: %w", fn, err)
	}

	if err = stmt.QueryRow(roleName, accessLevel).Scan(&res.Id, &res.RoleName, &res.AccessLeve); err != nil {
		return res, fmt.Errorf("%s: failed to get new created role: %w", fn, err)
	}

	return res, nil
}

// TODO: implement method:
// func (s *Storage) DeleteRole(id int) error
