package domain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string
	Users    []User
}

type RoleResponse struct {
	RoleID   uint                      `json:"id"`
	RoleName string                    `json:"name"`
	Users    []UserResponseWithOutRole `json:"users"`
}
