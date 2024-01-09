package service

import repository "github.com/lMikadal/POS_Backend_golang.git/repositories"

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) roleService {
	return roleService{roleRepo: roleRepo}
}

func (s roleService) GetRoles() ([]RoleResponse, error) {
	roles, err := s.roleRepo.GetAll()
	if err != nil {
		return nil, err
	}

	roleResponses := []RoleResponse{}
	for _, role := range roles {
		roleResponse := RoleResponse{
			RoleID:   role.ID,
			RoleName: role.RoleName,
		}
		roleResponses = append(roleResponses, roleResponse)
	}

	return roleResponses, nil
}

func (s roleService) GetRoleById(id int) (*RoleResponse, error) {
	role, err := s.roleRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	roleResponse := RoleResponse{
		RoleID:   role.ID,
		RoleName: role.RoleName,
	}
	return &roleResponse, nil
}
