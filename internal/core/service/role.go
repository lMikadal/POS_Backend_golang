package service

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type roleService struct {
	repo port.RoleRepository
}

func NewRoleService(repo port.RoleRepository) *roleService {
	return &roleService{
		repo,
	}
}

func (s roleService) GetRoles() ([]domain.RoleResponse, error) {
	roles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	roleResponses := []domain.RoleResponse{}
	for _, role := range roles {

		roleResponses = append(roleResponses, domain.RoleResponse{
			RoleID:   role.ID,
			RoleName: role.RoleName,
		})
	}

	return roleResponses, nil
}

func (s roleService) GetRoleById(id int) (*domain.RoleResponse, error) {
	role, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	roleResponse := domain.RoleResponse{
		RoleID:   role.ID,
		RoleName: role.RoleName,
	}
	return &roleResponse, nil
}
