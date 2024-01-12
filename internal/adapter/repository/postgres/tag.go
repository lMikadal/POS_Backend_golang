package repository

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

type tagRepository struct {
	db *DB
}

func NewTagRepository(db *DB) *tagRepository {
	return &tagRepository{db}
}

func (r tagRepository) GetAll() ([]domain.Tag, error) {
	tags := []domain.Tag{}

	err := r.db.Preload("Goodses").Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (r tagRepository) GetById(id int) (*domain.Tag, error) {
	tag := domain.Tag{}

	err := r.db.Preload("Goodses").First(&tag, id).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r tagRepository) Create(tag *domain.Tag) (*domain.Tag, error) {
	err := r.db.Create(&tag).Error
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (r tagRepository) Update(tag *domain.Tag, id int) (*domain.Tag, error) {
	newTag := domain.Tag{}

	err := r.db.First(&newTag, id).Error
	if err != nil {
		return nil, err
	}

	newTag.TagName = tag.TagName

	err = r.db.Save(&newTag).Error
	if err != nil {
		return nil, err
	}

	return &newTag, nil
}

func (r tagRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Tag{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
