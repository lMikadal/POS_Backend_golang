package service

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/util"
)

type goodsService struct {
	repo port.GoodsRepository
}

func NewGoodsService(repo port.GoodsRepository) *goodsService {
	return &goodsService{repo}
}

func (s goodsService) GetGoodses() ([]domain.GoodsResponse, error) {
	goodses, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	goodsResponses := []domain.GoodsResponse{}
	for _, goods := range goodses {
		goodsResponses = append(goodsResponses, util.ConverGoodsToGoodsResponse(&goods))
	}

	return goodsResponses, nil
}

func (s goodsService) GetGoodsById(id int) (*domain.GoodsResponse, error) {
	goods, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	goodsResponse := util.ConverGoodsToGoodsResponse(goods)

	return &goodsResponse, nil
}

func (s goodsService) CreateGoods(goods *domain.GoodsRequest) (*domain.GoodsResponse, error) {
	tags, err := s.repo.SearchTags(goods.Tags)
	if err != nil {
		return nil, err
	}

	goodsDomain := domain.Goods{
		GoodsName:   goods.GoodsName,
		GoodsCode:   goods.GoodsCode,
		GoodsAmount: goods.GoodsAmount,
		GoodsCost:   goods.GoodsCost,
		Tags:        tags,
	}

	goodsCreate, err := s.repo.Create(&goodsDomain)
	if err != nil {
		return nil, err
	}

	goodsResponse := util.ConverGoodsToGoodsResponse(goodsCreate)

	return &goodsResponse, nil
}

func (s goodsService) UpdateGoods(goods *domain.GoodsRequest, id int) (*domain.GoodsResponse, error) {
	tags, err := s.repo.SearchTags(goods.Tags)
	if err != nil {
		return nil, err
	}

	goodsDomain := domain.Goods{
		GoodsName:   goods.GoodsName,
		GoodsCode:   goods.GoodsCode,
		GoodsAmount: goods.GoodsAmount,
		GoodsCost:   goods.GoodsCost,
		Tags:        tags,
	}

	goodsUpdate, err := s.repo.Update(&goodsDomain, id)
	if err != nil {
		return nil, err
	}

	goodsResponse := util.ConverGoodsToGoodsResponse(goodsUpdate)

	return &goodsResponse, nil
}

func (s goodsService) DeleteGoods(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
