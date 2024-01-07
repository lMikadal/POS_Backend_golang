package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string
	UserEmail    string
	UserPassword string
	RoleID       uint
	Role         Role
}

type Role struct {
	gorm.Model
	RoleName string
	Users    []User
}

type UserResponse struct {
	UserID    uint         `json:"id"`
	UserName  string       `json:"name"`
	UserEmail string       `json:"email"`
	Role      RoleResponse `json:"role"`
}

type RoleResponse struct {
	RoleID   uint   `json:"id"`
	RoleName string `json:"role"`
}
