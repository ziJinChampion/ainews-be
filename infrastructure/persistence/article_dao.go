package persistence

import (
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/domain/repository"
	"gorm.io/gorm"
)

type ArticleDAO struct {
	db *gorm.DB
}

func (t *ArticleDAO) DeleteArticle(id int) error {
	if err := t.db.Delete(&entity.Article{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (t *ArticleDAO) GetArticles(m map[string]string) (article []*entity.Article, err error) {
	if err = t.db.Where(m).Find(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (t *ArticleDAO) CreateArticle(article *entity.Article, tagIds []int) (*entity.Article, error) {

	if err := t.db.Create(&article).Error; err != nil {
		return nil, err
	}

	for _, tagId := range tagIds {
		articleTag := &entity.ArticleTag{
			ArticleId: article.Id,
			TagId:     tagId,
		}
		if err := t.db.Create(&articleTag).Error; err != nil {
			return nil, err
		}
	}

	return article, nil
}

func (t *ArticleDAO) UpdateArticle(article *entity.Article) (*entity.Article, error) {
	if err := t.db.Save(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

var _ repository.ArticleRepository = &ArticleDAO{}

func NewArticleDAO(db *gorm.DB) *ArticleDAO {
	return &ArticleDAO{db}
}
