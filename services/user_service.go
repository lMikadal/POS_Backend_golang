package service

import (
	"log"

	repository "github.com/lMikadal/POS_Backend_golang.git/repositories"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			UserID:    user.ID,
			UserName:  user.UserName,
			UserEmail: user.UserEmail,
			Role: RoleResponse{
				RoleID:   user.RoleID,
				RoleName: user.Role.RoleName,
			},
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (s userService) GetUserById(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := UserResponse{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
	}
	return &userResponse, nil
}
