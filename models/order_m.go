package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderPriceTotal float32  `gorm:"type:decimal(10,2)"`
	User            []User   `gorm:"many2many:order_users;"`
	Status          []Status `gorm:"many2many:order_status;"`
	Goods           []Goods  `gorm:"many2many:order_goods;"`
}

type Status struct {
	gorm.Model
	StatusName string
}
