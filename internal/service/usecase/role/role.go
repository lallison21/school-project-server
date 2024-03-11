package usecase_role

import (
	"context"
	entity "github.com/lallison21/school-project-server/internal/entities"
)

type RoleRepository interface {
	GetRole(ctx context.Context, id int) (*entity.Role, error)
}

type RoleUseCase struct {
	roleRepository RoleRepository
}

func (repo *RoleUseCase) GetRole(ctx context.Context, id int) (*entity.Role, error) {
	return nil, nil
}
