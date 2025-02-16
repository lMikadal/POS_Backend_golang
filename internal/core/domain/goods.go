package domain

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsName   string
	GoodsCode   string
	GoodsAmount int
	GoodsCost   float32 `gorm:"type:decimal(10,2)"`
	Tags        []*Tag  `gorm:"many2many:goods_tags;"`
	Prices      []Price
}

type GoodsRequest struct {
	GoodsName   string  `json:"name"`
	GoodsCode   string  `json:"code"`
	GoodsAmount int     `json:"amount"`
	GoodsCost   float32 `json:"cost" gorm:"type:decimal(10,2)"`
	Tags        []int   `json:"tags"`
}

type GoodsResponse struct {
	GoodsID     uint                      `json:"id"`
	GoodsName   string                    `json:"name"`
	GoodsCode   string                    `json:"code"`
	GoodsAmount int                       `json:"amount"`
	GoodsCost   float32                   `json:"cost" gorm:"type:decimal(10,2)"`
	Tags        []TagResponseWithOutGoods `json:"tags"`
	Prices      []PriceResponse           `json:"prices"`
}

type GoodsResponseWithOutTag struct {
	GoodsID     uint    `json:"id"`
	GoodsName   string  `json:"name"`
	GoodsCode   string  `json:"code"`
	GoodsAmount int     `json:"amount"`
	GoodsCost   float32 `json:"cost" gorm:"type:decimal(10,2)"`
}
