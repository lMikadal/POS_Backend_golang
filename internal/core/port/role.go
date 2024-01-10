package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type RoleRepository interface {
	GetAll() ([]domain.Role, error)
	GetById(int) (*domain.Role, error)
}

type RoleService interface {
	GetRoles() ([]domain.RoleResponse, error)
	GetRoleById(int) (*domain.RoleResponse, error)
}
