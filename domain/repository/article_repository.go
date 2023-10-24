package repository

import "github.com/southwind/ainews/domain/entity"

type ArticleRepository interface {
	CreateArticle(*entity.Article, []int) (*entity.Article, error)
	GetArticles(map[string]string) ([]*entity.Article, error)
	UpdateArticle(*entity.Article) (*entity.Article, error)
	DeleteArticle(id int) error
}
