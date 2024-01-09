package repository

import "gorm.io/gorm"

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll() ([]User, error) {
	users := []User{}

	err := r.db.Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepositoryDB) GetById(id int) (*User, error) {
	user := User{}

	err := r.db.Preload("Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
