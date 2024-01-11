package service

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/port"

type goodsService struct {
	repo port.GoodsRepository
}

func NewGoodsService(repo port.GoodsRepository) *goodsService {
	return &goodsService{repo}
}
