package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type PriceRepository interface {
	Create(*domain.Price) (*domain.Price, error)
	Update(*domain.Price, int) (*domain.Price, error)
	Delete(int) error
}
