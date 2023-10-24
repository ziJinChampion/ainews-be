package application

import (
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/domain/repository"
)

type articleApp struct {
	articleRepo repository.ArticleRepository
}

type ArticleInterface interface {
	CreateArticle(*entity.Article, []int) (*entity.Article, error)
	GetArticles(map[string]string) ([]*entity.Article, error)
	UpdateArticle(*entity.Article) (*entity.Article, error)
	DeleteArticle(id int) error
}

func (t *articleApp) GetArticles(m map[string]string) ([]*entity.Article, error) {
	return t.articleRepo.GetArticles(m)
}

func (t *articleApp) UpdateArticle(article *entity.Article) (*entity.Article, error) {
	return t.articleRepo.UpdateArticle(article)
}

func (t *articleApp) CreateArticle(article *entity.Article, tagIds []int) (*entity.Article, error) {
	return t.articleRepo.CreateArticle(article, tagIds)
}

func (t *articleApp) DeleteArticle(id int) error {
	return t.articleRepo.DeleteArticle(id)
}

var _ ArticleInterface = &articleApp{}
