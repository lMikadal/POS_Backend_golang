package repository

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
)

type goodsRepository struct {
	db *DB
}

func NewGoodsRepository(db *DB) *goodsRepository {
	return &goodsRepository{db}
}

func (r goodsRepository) GetAll() ([]domain.Goods, error) {
	goodses := []domain.Goods{}

	err := r.db.Preload("Tags").Preload("Prices").Find(&goodses).Error
	if err != nil {
		return nil, err
	}

	return goodses, nil
}

func (r goodsRepository) GetById(id int) (*domain.Goods, error) {
	goods := domain.Goods{}

	err := r.db.Preload("Tags").Preload("Prices").First(&goods, id).Error
	if err != nil {
		return nil, err
	}

	return &goods, nil
}

func (r goodsRepository) Create(goods *domain.Goods) (*domain.Goods, error) {
	err := r.db.Create(&goods).Error
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func (r goodsRepository) Update(goods *domain.Goods, id int) (*domain.Goods, error) {
	var goodsDB domain.Goods
	err := r.db.Preload("Tags").First(&goodsDB, id).Error
	if err != nil {
		return nil, err
	}

	goodsDB.GoodsName = goods.GoodsName
	goodsDB.GoodsCode = goods.GoodsCode
	goodsDB.GoodsAmount = goods.GoodsAmount
	goodsDB.GoodsCost = goods.GoodsCost

	err = r.db.Model(&goodsDB).Association("Tags").Clear()
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&goodsDB).Association("Tags").Append(goods.Tags)
	if err != nil {
		return nil, err
	}

	err = r.db.Save(&goodsDB).Error
	if err != nil {
		return nil, err
	}

	return &goodsDB, nil
}

func (r goodsRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Goods{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r goodsRepository) SearchTags(tags []int) ([]*domain.Tag, error) {
	tagDomain := []*domain.Tag{}
	for _, tagId := range tags {
		tag := domain.Tag{}
		err := r.db.First(&tag, tagId).Error
		if err != nil {
			return nil, err
		}
		tagDomain = append(tagDomain, &tag)
	}
	return tagDomain, nil
}
