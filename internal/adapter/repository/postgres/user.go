package repository

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
)

type userRepository struct {
	db *DB
}

func NewUserRepository(db *DB) *userRepository {
	return &userRepository{db}
}

func (r userRepository) GetAll() ([]domain.User, error) {
	users := []domain.User{}

	err := r.db.Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepository) GetById(id int) (*domain.User, error) {
	user := domain.User{}

	err := r.db.Preload("Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
