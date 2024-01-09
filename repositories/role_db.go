package repository

import "gorm.io/gorm"

type roleRepositoryDB struct {
	db *gorm.DB
}

func NewRoleRepositoryDB(db *gorm.DB) roleRepositoryDB {
	return roleRepositoryDB{db: db}
}

func (r roleRepositoryDB) GetAll() ([]Role, error) {
	roles := []Role{}

	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r roleRepositoryDB) GetById(id int) (*Role, error) {
	role := Role{}

	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}
