package domain

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName string
	Goodses []*Goods `gorm:"many2many:goods_tags;"`
}
