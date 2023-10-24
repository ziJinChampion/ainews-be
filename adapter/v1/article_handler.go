package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/application"
	"github.com/southwind/ainews/common/code"
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/utils/security"
)

type Articles struct {
	articleApplcation application.ArticleInterface
	tagApplication    application.TagAppInterface
}

func NewArticles(articleApplcation application.ArticleInterface,
	tagApplication application.TagAppInterface) *Articles {
	return &Articles{articleApplcation, tagApplication}
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
