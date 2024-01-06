package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	StockAmount uint16
	StockCost   float32 `gorm:"type:decimal(10,2)"`
	StockDate   time.Time
	Goods       []Goods `gorm:"many2many:stock_goods;"`
}
