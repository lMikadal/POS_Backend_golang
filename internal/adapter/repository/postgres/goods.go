package repository

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type goodsRepository struct {
	db *DB
}

func NewGoodsRepository(db *DB) *goodsRepository {
	return &goodsRepository{db}
}

func (r goodsRepository) GetAll() ([]domain.Goods, error) {
	return nil, nil
}

func (r goodsRepository) GetById(id int) (*domain.Goods, error) {
	return nil, nil
}

func (r goodsRepository) Create(*domain.Goods) (*domain.Goods, error) {
	return nil, nil
}

func (r goodsRepository) Update(*domain.Goods) (*domain.Goods, error) {
	return nil, nil
}

func (r goodsRepository) Delete(id int) error {
	return nil
}
