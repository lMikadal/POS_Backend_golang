package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type TagRepository interface {
	GetAll() ([]domain.Tag, error)
	GetById(int) (*domain.Tag, error)
	Create(*domain.Tag) (*domain.Tag, error)
	Update(*domain.Tag) (*domain.Tag, error)
	Delete(int) error
}
