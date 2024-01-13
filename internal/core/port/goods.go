package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type GoodsRepository interface {
	GetAll() ([]domain.Goods, error)
	GetById(int) (*domain.Goods, error)
	Create(*domain.Goods) (*domain.Goods, error)
	Update(*domain.Goods, int) (*domain.Goods, error)
	Delete(int) error
	SearchTags([]int) ([]*domain.Tag, error)
}

type GoodsService interface {
	GetGoodses() ([]domain.GoodsResponse, error)
	GetGoodsById(int) (*domain.GoodsResponse, error)
	CreateGoods(*domain.GoodsRequest) (*domain.GoodsResponse, error)
	UpdateGoods(*domain.GoodsRequest, int) (*domain.GoodsResponse, error)
	DeleteGoods(int) error
}
