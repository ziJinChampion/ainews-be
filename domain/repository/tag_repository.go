package repository

import "github.com/southwind/ainews/domain/entity"

type TagRepository interface {
	CreateTag(*entity.Tag) (*entity.Tag, error)
	GetTags(map[string]interface{}) ([]*entity.Tag, error)
	UpdateTag(*entity.Tag) (*entity.Tag, error)
	DeleteTag(id int) error
}
