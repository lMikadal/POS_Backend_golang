package service

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type priceService struct {
	repo port.PriceRepository
}

func NewPriceService(repo port.PriceRepository) *priceService {
	return &priceService{repo}
}

func (s priceService) CreatePrice(price *domain.PriceRequest) (*domain.PriceResponse, error) {
	priceDB := &domain.Price{
		PriceAmount: price.PriceAmount,
		PricePrice:  price.PricePrice,
		GoodsID:     price.GoodsID,
	}

	priceCreate, err := s.repo.Create(priceDB)
	if err != nil {
		return nil, err
	}

	priceResponse := &domain.PriceResponse{
		PriceID:     priceCreate.ID,
		PriceAmount: priceCreate.PriceAmount,
		PricePrice:  priceCreate.PricePrice,
	}

	return priceResponse, nil
}

func (s priceService) UpdatePrice(price *domain.PriceRequest, id int) (*domain.PriceResponse, error) {
	priceDB := &domain.Price{
		PriceAmount: price.PriceAmount,
		PricePrice:  price.PricePrice,
		GoodsID:     price.GoodsID,
	}

	priceUpdate, err := s.repo.Update(priceDB, id)
	if err != nil {
		return nil, err
	}

	priceResponse := &domain.PriceResponse{
		PriceID:     priceUpdate.ID,
		PriceAmount: priceUpdate.PriceAmount,
		PricePrice:  priceUpdate.PricePrice,
	}

	return priceResponse, nil
}

func (s priceService) DeletePrice(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
