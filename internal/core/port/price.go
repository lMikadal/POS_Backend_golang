package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type PriceRepository interface {
	Create(*domain.Price) (*domain.Price, error)
	Update(*domain.Price, int) (*domain.Price, error)
	Delete(int) error
}

type PriceService interface {
	CreatePrice(*domain.Price) (*domain.Price, error)
	UpdatePrice(*domain.Price, int) (*domain.Price, error)
	DeletePrice(int) error
}
