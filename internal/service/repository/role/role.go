package repository_role

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lallison21/school-project-server/internal/entities"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetRole(ctx context.Context, id int) (*entity.Role, error) {
	return nil, nil
}
