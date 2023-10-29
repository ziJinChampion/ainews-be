package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/application"
	"github.com/southwind/ainews/common/code"
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/utils/security"
)

type Articles struct {
	articleApplcation application.ArticleInterface
	tagApplication    application.TagAppInterface
	userApplication   application.UserAppInterface
}

func NewArticles(articleApplcation application.ArticleInterface,
	tagApplication application.TagAppInterface,
	userApplication application.UserAppInterface) *Articles {
	return &Articles{articleApplcation, tagApplication, userApplication}
}

type CreateArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Tags    []int  `json:"tags" binding:"required"`
}

func (t *Articles) CreateArticle(c *gin.Context) {
	var req CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}
	userInfo, err := security.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code.NO_AUTH,
			"msg":  err.Error(),
		})
		return
	}

	var article = &entity.Article{
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: userInfo.Id,
	}

	if tags, _ := t.tagApplication.GetTags(map[string]interface{}{"id": req.Tags}); len(tags) < len(req.Tags) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  "tag not exist",
		})
		return
	}

	if _, err := article.VaildContentInfo(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}
	article.PrepareCreate()
	res, err := t.articleApplcation.CreateArticle(article, req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code.ERROR,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"data": res,
	})
}

type ArticleVO struct {
	Id         int           `json:"id" gorm:"primary_key;autoIncrement"`
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	AuthorName string        `json:"author_name"`
	Tags       []*entity.Tag `json:"tags"`
}

func (t *Articles) GetArticles(c *gin.Context) {
	var id = c.Query("id")
	var pageSize = c.Query("pageSize")
	var pageNum = c.Query("pageNum")
	if pageSize == "" {
		pageSize = "10"
	}
	if pageNum == "" {
		pageNum = "1"
	}
	var queryMap = make(map[string]interface{})
	if id != "" {
		queryMap["id"] = id
	}
	var res []*ArticleVO
	size, _ := strconv.Atoi(pageSize)
	num, _ := strconv.Atoi(pageNum)
	articles, err := t.articleApplcation.GetArticles(queryMap, size, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}
	for _, article := range articles {
		vo := &ArticleVO{
			Id:        article.Id,
			Title:     article.Title,
			Content:   article.Content,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		}
		user, _ := t.userApplication.GetUsers(map[string]interface{}{"id": article.AuthorId})
		if len(user) == 1 {
			vo.AuthorName = user[0].Name
		}
		articleTags, _ := t.articleApplcation.GetArticleTags(article.Id)
		tagIds := make([]int, len(articleTags))
		for i, articleTag := range articleTags {
			tagIds[i] = articleTag.TagId
		}
		tags, _ := t.tagApplication.GetTags(map[string]interface{}{"id": tagIds})
		if len(tags) > 0 {
			vo.Tags = tags
		}
		res = append(res, vo)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"data": res,
	})

}
