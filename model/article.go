package model

import (
	"errors"

	"gorm.io/gorm"
)

type Article struct {
	Model

	Title    string `gorm:"column:title" json:"article_title"`
	Content  string `gorm:"column:content" json:"article_content"`
	AuthorId int    `gorm:"column:author_id" json:"article_author"`
	IsHidden bool   `gorm:"column:is_hidden" json:"is_hidden"`
}

type ArticleResponse struct {
	Id        int      `json:"id"`
	Title     string   `json:"article_title"`
	Content   string   `json:"article_content"`
	Author    string   `json:"article_author"`
	Tags      []string `json:"tags"`
	UpdatedAt string   `json:"updated_at"`
}

type ArticleTag struct {
	Model

	ArticleId int `gorm:"column:article_id"`
	TagId     int `gorm:"column:tag_id"`
}

func GetAllArticles() (articles []ArticleResponse, err error) {
	err = client.Table("articles").Select("articles.id, articles.title, articles.content, articles.updated_at, users.user_name as author, ARRAY_AGG(t.tag_name) as tags").
		Joins("left join users on users.id = articles.author_id").
		Joins("left join article_tag on articles.id = article_tag.article_id").
		Joins("left join tags t on article_tag.tag_id = t.id").
		Group("articles.id").
		Group("users.user_name").Scan(&articles).Error
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

		if err = tx.Table("articles").Create(&article).Error; err != nil {
			return err
		}
		for _, tagId := range tagIds {
			var tag Tag
			if err = tx.Table("tags").Where("id = ?", tagId).First(&tag).Error; err != nil {
				return errors.New("tag not found")
			}
			if err = tx.Table("article_tag").Create(&ArticleTag{
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
