package domain

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsName   string
	GoodsCode   string
	GoodsAmount int
	GoodsCost   float32
	Tags        []*Tag `gorm:"many2many:goods_tags;"`
}
