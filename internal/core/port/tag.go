package port

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type TagRepository interface {
	GetAll() ([]domain.Tag, error)
	GetById(int) (*domain.Tag, error)
	Create(*domain.Tag) (*domain.Tag, error)
	Update(*domain.Tag, int) (*domain.Tag, error)
	Delete(int) error
}

type TagService interface {
	GetTags() ([]domain.TagResponse, error)
	GetTagById(int) (*domain.TagResponse, error)
	CreateTag(*domain.TagRequest) (*domain.TagResponse, error)
	UpdateTag(*domain.TagRequest, int) (*domain.TagResponse, error)
	DeleteTag(int) error
}
