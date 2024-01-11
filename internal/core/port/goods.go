package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type GoodsRepository interface {
	GetAll() ([]domain.Goods, error)
	GetById(int) (*domain.Goods, error)
	Create(*domain.Goods) (*domain.Goods, error)
	Update(*domain.Goods) (*domain.Goods, error)
	Delete(int) error
}

type GoodsService interface {
	GetGoodses() ([]domain.GoodsResponse, error)
	GetGoodsById(int) (*domain.GoodsResponse, error)
	CreateGoods(*domain.GoodsResponse) (*domain.GoodsResponse, error)
	UpdateGoods(*domain.GoodsResponse) (*domain.GoodsResponse, error)
	DeleteGoods(int) error
}
