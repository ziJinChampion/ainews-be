package model

import "gorm.io/gorm"

type Article struct {
	Model

	Title    string `gorm:"column:title" json:"article_title"`
	Content  string `gorm:"column:content" json:"article_content"`
	AuthorId int    `gorm:"column:author_id" json:"article_author"`
	TagId    int    `gorm:"column:tag_id" json:"tag_id"`
	IsHidden bool   `gorm:"column:is_hidden" json:"is_hidden"`
}

type ArticleVO struct {
	Id         int      `json:"id"`
	Title      string   `json:"article_title"`
	Content    string   `json:"article_content"`
	AuthorName string   `json:"article_author"`
	Tags       []string `json:"tags"`
	IsHidden   bool     `json:"is_hidden"`
}

type ArticleTag struct {
	Model

	ArticleId int `gorm:"column:article_id"`
	TagId     int `gorm:"column:tag_id"`
}

func GetAllArticles() (articles []Article, err error) {
	err = client.Table("articles").Find(&articles).Error
	return
}

func GetArticlesByTag(tagId int) (articles []Article, err error) {
	err = client.Table("articles").Where("tag_id = ?", tagId).Find(&articles).Error
	return
}

func GetArticlesByAuthor(authorId int) (articles []Article, err error) {
	err = client.Table("articles").Where("article_author = ?", authorId).Find(&articles).Error
	return
}

func CreateNewArticle(title, content string, authorId int, tagIds []int) (bool, error) {

	var err error

	client.Transaction(func(tx *gorm.DB) error {

		article := Article{
			Title:    title,
			Content:  content,
			AuthorId: authorId,
			IsHidden: false}

		if err = client.Table("articles").Create(&article).Error; err != nil {
			return err
		}
		for _, tagId := range tagIds {
			if err = client.Table("article_tag").Create(&ArticleTag{
				ArticleId: article.Id,
				TagId:     tagId,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return true, err
}
