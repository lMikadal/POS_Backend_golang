package domain

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	PriceAmount int
	PricePrice  float32 `gorm:"type:decimal(10,2)"`
	GoodsID     uint
}

type PriceRequest struct {
	PriceAmount int     `json:"amount"`
	PricePrice  float32 `json:"price"`
	GoodsID     uint    `json:"goods_id"`
}

type PriceResponse struct {
	PriceID     uint    `json:"id"`
	PriceAmount int     `json:"amount"`
	PricePrice  float32 `json:"price"`
}
