package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type PriceRepository interface {
	Create(*domain.Price) (*domain.Price, error)
	Update(*domain.Price, int) (*domain.Price, error)
	Delete(int) error
}

type PriceService interface {
	CreatePrice(*domain.PriceRequest) (*domain.PriceResponse, error)
	UpdatePrice(*domain.PriceRequest, int) (*domain.PriceResponse, error)
	DeletePrice(int) error
}
