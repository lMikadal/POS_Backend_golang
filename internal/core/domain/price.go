package domain

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	PriceAmount int
	PriceCost   float32 `gorm:"type:decimal(10,2)"`
	GoodsID     uint
}

type PriceResponse struct {
	PriceID     uint    `json:"id"`
	PriceAmount int     `json:"amount"`
	PriceCost   float32 `json:"cost"`
}
