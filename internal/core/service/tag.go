package service

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/port"
)

type tagService struct {
	repo port.TagRepository
}

func NewTagService(repo port.TagRepository) *tagService {
	return &tagService{repo}
}

func (s tagService) GetTags() ([]domain.TagResponse, error) {
	tags, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	tagsResponse := []domain.TagResponse{}
	for _, tag := range tags {
		tagsResponse = append(tagsResponse, domain.TagResponse{
			TagID:   tag.ID,
			TagName: tag.TagName,
		})
	}

	return tagsResponse, nil
}

func (s tagService) GetTagById(id int) (*domain.TagResponse, error) {
	tag, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	tagResponse := domain.TagResponse{
		TagID:   tag.ID,
		TagName: tag.TagName,
	}

	return &tagResponse, nil
}

func (s tagService) CreateTag(tag *domain.TagRequest) (*domain.TagResponse, error) {
	tagDomain := domain.Tag{
		TagName: tag.TagName,
	}

	tagCreate, err := s.repo.Create(&tagDomain)
	if err != nil {
		return nil, err
	}

	tagResponse := domain.TagResponse{
		TagID:   tagCreate.ID,
		TagName: tagCreate.TagName,
	}

	return &tagResponse, nil
}

func (s tagService) UpdateTag(tag *domain.TagRequest, id int) (*domain.TagResponse, error) {
	tagDomain := domain.Tag{
		TagName: tag.TagName,
	}

	tagUpdate, err := s.repo.Update(&tagDomain, id)
	if err != nil {
		return nil, err
	}

	tagResponse := domain.TagResponse{
		TagID:   tagUpdate.ID,
		TagName: tagUpdate.TagName,
	}

	return &tagResponse, nil
}

func (s tagService) DeleteTag(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
