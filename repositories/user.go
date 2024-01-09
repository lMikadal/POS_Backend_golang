package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string
	UserEmail    string
	UserPassword string
	RoleID       uint
	Role         Role
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(int) (*User, error)
}
