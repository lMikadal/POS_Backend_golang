package models

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsCode   string
	GoodsName   string
	GoodsAmount uint16
	GoodsCost   float32 `gorm:"type:decimal(10,2)"`
	GoodsPrice  []GoodsPrice
	Tag         []Tag `gorm:"many2many:goods_tags;"`
}

type GoodsPrice struct {
	gorm.Model
	GoodsPriceAmount uint16
	GoodsPricePrice  float32 `gorm:"type:decimal(10,2)"`
	GoodsID          uint
}

type Tag struct {
	gorm.Model
	TagName string
}
