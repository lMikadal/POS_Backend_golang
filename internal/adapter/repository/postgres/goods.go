package repository

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type goodsRepository struct {
	db *DB
}

func NewGoodsRepository(db *DB) *goodsRepository {
	return &goodsRepository{db}
}

func (r goodsRepository) GetAll() ([]domain.Goods, error) {
	goodses := []domain.Goods{}

	err := r.db.Preload("Tags").Find(&goodses).Error
	if err != nil {
		return nil, err
	}

	return goodses, nil
}

func (r goodsRepository) GetById(id int) (*domain.Goods, error) {
	goods := domain.Goods{}

	err := r.db.Preload("Tags").First(&goods, id).Error
	if err != nil {
		return nil, err
	}

	return &goods, nil
}

func (r goodsRepository) Create(goods *domain.Goods) (*domain.Goods, error) {
	err := r.db.Create(goods).Error
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func (r goodsRepository) Update(goods *domain.Goods) (*domain.Goods, error) {
	err := r.db.Save(goods).Error
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func (r goodsRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Goods{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
