package service

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *userService {
	return &userService{
		repo,
	}
}

func (s userService) GetUsers() ([]domain.UserResponse, error) {
	users, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	userResponses := []domain.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, domain.UserResponse{
			UserID:    user.ID,
			UserName:  user.UserName,
			UserEmail: user.UserEmail,
			Role: domain.RoleResponse{
				RoleID:   user.RoleID,
				RoleName: user.Role.RoleName,
			},
		})
	}

	return userResponses, nil
}

func (s userService) GetUserById(id int) (*domain.UserResponse, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	userResponse := domain.UserResponse{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		Role: domain.RoleResponse{
			RoleID:   user.RoleID,
			RoleName: user.Role.RoleName,
		},
	}
	return &userResponse, nil
}
