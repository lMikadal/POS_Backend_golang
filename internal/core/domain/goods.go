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

type GoodsRequest struct {
	GoodsId     uint    `json:"id"`
	GoodsName   string  `json:"name"`
	GoodsCode   string  `json:"code"`
	GoodsAmount int     `json:"amount"`
	GoodsCost   float32 `json:"cost"`
	Tags        []uint  `json:"tags"`
}

type GoodsResponse struct {
	GoodsID     uint                      `json:"id"`
	GoodsName   string                    `json:"name"`
	GoodsCode   string                    `json:"code"`
	GoodsAmount int                       `json:"amount"`
	GoodsCost   float32                   `json:"cost"`
	Tags        []TagResponseWithOutGoods `json:"tags"`
}

type GoodsResponseWithOutTag struct {
	GoodsID     uint    `json:"id"`
	GoodsName   string  `json:"name"`
	GoodsCode   string  `json:"code"`
	GoodsAmount int     `json:"amount"`
	GoodsCost   float32 `json:"cost"`
}
