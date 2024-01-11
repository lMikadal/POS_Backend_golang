package domain

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName string
	Goodses []*Goods `gorm:"many2many:goods_tags;"`
}

type TagResponse struct {
	TagID   uint                      `json:"id"`
	TagName string                    `json:"name"`
	Goodses []GoodsResponseWithOutTag `json:"goodses"`
}

type TagResponseWithOutGoods struct {
	TagID   uint   `json:"id"`
	TagName string `json:"name"`
}
