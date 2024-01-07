package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string
	UserEmail    string
	UserPassword string
	Role         []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	gorm.Model
	RoleName string
}

type UserResponse struct {
	UserName  string `json:"name"`
	UserEmail string `json:"email"`
	RoleName  string `json:"role"`
}
