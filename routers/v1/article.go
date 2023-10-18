package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/model"
	"github.com/southwind/ainews/pkg/e"
	"github.com/southwind/ainews/pkg/utils"
)

type CreateArticleForm struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	TagIds  []int  `json:"tag_ids" binding:"required"`
}

func CreateNewArticle(c *gin.Context) {
	var form CreateArticleForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  err.Error()})
		return
	}

	vaild := validation.Validation{}
	vaild.Required(form.Title, "title").Message("Must fill Article Title")
	vaild.Required(form.Content, "content").Message("Must fill Article Content")
	vaild.Required(form.TagIds, "tag_ids").Message("Must fill Article Tag")
	vaild.MaxSize(form.Title, 80, "title").Message("Article Title can not over 80 character")

	if vaild.HasErrors() {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  vaild.Errors[0].Message})
		return
	}

	if currentUser, err := utils.GetCurrentUser(c); err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"code": e.ERROR_AUTH_CHECK_TOKEN_FAIL,
			"msg":  e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL)})
		return
	} else {
		if _, err := model.CreateNewArticle(form.Title, form.Content, currentUser.Id, form.TagIds); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": e.ERROR,
				"msg":  "Create Article Fail"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": e.SUCCESS,
				"msg":  "Create Article Success"})
		}
	}

}

type ArticleResponse struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	Author    string   `json:"author"`
	UpdatedAt string   `json:"updated_at"`
}

func GetAllArticles(c *gin.Context) {
	if articles, err := model.GetAllArticles(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": e.ERROR,
			"msg":  err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": articles})
	}
}
