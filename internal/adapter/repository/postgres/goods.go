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

	err := r.db.Preload("Tag").Find(&goodses).Error
	if err != nil {
		return nil, err
	}

	return goodses, nil
}

func (r goodsRepository) GetById(id int) (*domain.Goods, error) {
	goods := domain.Goods{}

	err := r.db.Preload("Tag").First(&goods, id).Error
	if err != nil {
		return nil, err
	}

	return &goods, nil
}

func (r goodsRepository) Create(goods *domain.GoodsRequest) (*domain.Goods, error) {
	tagDB := []*domain.Tag{}
	for _, tagId := range goods.Tags {
		tag := domain.Tag{}
		err := r.db.First(&tag, tagId).Error
		if err != nil {
			return nil, err
		}
		tagDB = append(tagDB, &tag)
	}

	goodsCreate := domain.Goods{
		GoodsName:   goods.GoodsName,
		GoodsCode:   goods.GoodsCode,
		GoodsAmount: goods.GoodsAmount,
		GoodsCost:   goods.GoodsCost,
		Tags:        tagDB,
	}

	err := r.db.Create(&goodsCreate).Error
	if err != nil {
		return nil, err
	}

	return &goodsCreate, nil
}

func (r goodsRepository) Update(goods *domain.GoodsRequest, id int) (*domain.Goods, error) {
	goodsDB := domain.Goods{}
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
	tagDB := []domain.Tag{}
	for _, tagId := range goods.Tags {
		tag := domain.Tag{}
		err := r.db.First(&tag, tagId).Error
		if err != nil {
			return nil, err
		}
		tagDB = append(tagDB, tag)
	}
	r.db.Model(&goodsDB).Association("Tags").Append(tagDB)

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
