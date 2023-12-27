package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Test string `json:"test"`
}
