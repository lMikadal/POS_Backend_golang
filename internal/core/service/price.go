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

func (s priceService) CreatePrice(price *domain.Price) (*domain.Price, error) {
	priceCreate, err := s.repo.Create(price)
	if err != nil {
		return nil, err
	}

	return priceCreate, nil
}

func (s priceService) UpdatePrice(price *domain.Price, id int) (*domain.Price, error) {
	priceUpdate, err := s.repo.Update(price, id)
	if err != nil {
		return nil, err
	}

	return priceUpdate, nil
}

func (s priceService) DeletePrice(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
