package service

type RoleResponse struct {
	RoleID   uint   `json:"id"`
	RoleName string `json:"name"`
}

type RoleService interface {
	GetRoles() ([]RoleResponse, error)
	GetRoleById(int) (*RoleResponse, error)
}
