package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type GoodsRepository interface {
	GetAll() ([]domain.Goods, error)
	GetById(int) (*domain.Goods, error)
	Create(*domain.GoodsRequest) (*domain.Goods, error)
	Update(*domain.GoodsRequest, int) (*domain.Goods, error)
	Delete(int) error
}

type GoodsService interface {
	GetGoodses() ([]domain.GoodsResponse, error)
	GetGoodsById(int) (*domain.GoodsResponse, error)
	CreateGoods(*domain.GoodsRequest) (*domain.GoodsResponse, error)
	UpdateGoods(*domain.GoodsRequest) (*domain.GoodsResponse, error)
	DeleteGoods(int) error
}
