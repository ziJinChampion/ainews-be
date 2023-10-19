package application

import (
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/domain/repository"
)

type tagApp struct {
	tagRepo repository.TagRepository
}

type TagAppInterface interface {
	CreateTag(*entity.Tag) (*entity.Tag, error)
	GetTags(map[string]string) ([]*entity.Tag, error)
	UpdateTag(*entity.Tag) (*entity.Tag, error)
	DeleteTag(id int) error
}

var _ TagAppInterface = &tagApp{}

func (t *tagApp) GetTags(m map[string]string) ([]*entity.Tag, error) {
	return t.tagRepo.GetTags(m)
}

func (t *tagApp) UpdateTag(tag *entity.Tag) (*entity.Tag, error) {
	return t.tagRepo.UpdateTag(tag)
}

func (t *tagApp) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	return t.tagRepo.CreateTag(tag)
}

func (t *tagApp) DeleteTag(id int) error {
	return t.tagRepo.DeleteTag(id)
}
