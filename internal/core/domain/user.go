package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string
	UserEmail    string
	UserPassword string
	RoleID       uint
	Role         Role
}

type UserResponse struct {
	UserID    uint                    `json:"id"`
	UserName  string                  `json:"name"`
	UserEmail string                  `json:"email"`
	Role      RoleResponseWithOutUser `json:"role"`
}

type UserResponseWithOutRole struct {
	UserID    uint   `json:"id"`
	UserName  string `json:"name"`
	UserEmail string `json:"email"`
}

func (UserResponseWithOutRole) TableName() string {
	return "users"
}
