package repository

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type tagRepository struct {
	db *DB
}

func NewTagRepository(db *DB) *tagRepository {
	return &tagRepository{db}
}

func (r tagRepository) GetAll() ([]domain.Tag, error) {
	return nil, nil
}

func (r tagRepository) GetById(id int) (*domain.Tag, error) {
	return nil, nil
}

func (r tagRepository) Create(*domain.Tag) (*domain.Tag, error) {
	return nil, nil
}

func (r tagRepository) Update(*domain.Tag) (*domain.Tag, error) {
	return nil, nil
}

func (r tagRepository) Delete(int) error {
	return nil
}
