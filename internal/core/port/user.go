package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetById(int) (*domain.User, error)
}

type UserService interface {
	GetUsers() ([]domain.UserResponse, error)
	GetUserById(int) (*domain.UserResponse, error)
}
