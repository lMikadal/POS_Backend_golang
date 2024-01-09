package service

type UserResponse struct {
	UserID    uint         `json:"id"`
	UserName  string       `json:"name"`
	UserEmail string       `json:"email"`
	Role      RoleResponse `json:"role"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUserById(int) (*UserResponse, error)
}
