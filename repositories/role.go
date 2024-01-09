package repository

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string
	Users    []User
}

type RoleRepository interface {
	GetAll() ([]Role, error)
	GetById(int) (*Role, error)
}
