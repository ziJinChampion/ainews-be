package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/application"
	"github.com/southwind/ainews/common/code"
	"github.com/southwind/ainews/domain/entity"
	"gorm.io/gorm"
)

type Tags struct {
	tagApplication application.TagAppInterface
}

func NewTag(tagApp application.TagAppInterface) *Tags {
	return &Tags{tagApp}
}

type CreateTagRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *Tags) CreateTag(c *gin.Context) {
	var req CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}
	var queryMap = make(map[string]string)
	queryMap["name"] = req.Name
	var tags []*entity.Tag
	var err error
	tags, err = t.tagApplication.GetTags(queryMap)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
	}
	if len(tags) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  "标签已存在",
		})
		return
	}
	tag, err := t.tagApplication.CreateTag(&entity.Tag{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code.ERROR,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code.SUCCESS,
			"data": tag.PublicTag(),
		})
	}

}

func (t *Tags) GetTags(c *gin.Context) {
	name := c.Param("name")
	var queryMap = make(map[string]string)
	if name != "" {
		queryMap["name"] = name
	}
	var tags []*entity.Tag
	var err error
	tags, err = t.tagApplication.GetTags(queryMap)
	var publicTags []*entity.PubicTag
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code.ERROR,
			"msg":  err.Error(),
		})
	} else {
		for _, tag := range tags {
			publicTags = append(publicTags, tag.PublicTag())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code.SUCCESS,
			"data": publicTags,
		})
	}

}

func (t *Tags) DeleteTag(c *gin.Context) {
	pathId := c.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
	}

	if err := t.tagApplication.DeleteTag(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"code": code.ERROR,
				"msg":  "tag not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": code.ERROR,
				"msg":  err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code.SUCCESS,
		})
	}
}
