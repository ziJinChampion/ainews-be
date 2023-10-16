package model

type Article struct {
	Model

	Title    string `gorm:"column:title" json:"article_title"`
	Content  string `gorm:"column:content" json:"article_content"`
	AuthorId int    `gorm:"column:author_id" json:"article_author"`
	TagId    int    `gorm:"column:tag_id" json:"tag_id"`
	IsHidden bool   `gorm:"column:is_hidden" json:"is_hidden"`
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

func CreateNewArticle(title, content string, authorId, tagId int) (bool, error) {
	err := client.Table("articles").Create(&Article{
		Title:    title,
		Content:  content,
		AuthorId: authorId,
		TagId:    tagId,
		IsHidden: false,
	}).Error

	return true, err
}
