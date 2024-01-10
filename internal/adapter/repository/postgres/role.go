package repository

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type roleRepository struct {
	db *DB
}

func NewRoleRepository(db *DB) *roleRepository {
	return &roleRepository{db}
}

func (r roleRepository) GetAll() ([]domain.Role, error) {
	roles := []domain.Role{}

	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r roleRepository) GetById(id int) (*domain.Role, error) {
	role := domain.Role{}

	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}
