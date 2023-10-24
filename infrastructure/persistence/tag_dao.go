package persistence

import (
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/domain/repository"
	"gorm.io/gorm"
)

type TagDAO struct {
	db *gorm.DB
}

func (t *TagDAO) DeleteTag(id int) error {
	if err := t.db.Delete(&entity.Tag{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (t *TagDAO) GetTags(m map[string]interface{}) (tag []*entity.Tag, err error) {
	if err = t.db.Where(m).Find(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (t *TagDAO) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	if err := t.db.Create(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (t *TagDAO) UpdateTag(tag *entity.Tag) (*entity.Tag, error) {
	if err := t.db.Save(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

var _ repository.TagRepository = &TagDAO{}

func NewTagDAO(db *gorm.DB) *TagDAO {
	return &TagDAO{db}
}
