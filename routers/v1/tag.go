package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/model"
	"github.com/southwind/ainews/pkg/e"
)

func GetAllTags(c *gin.Context) {

	if tags, err := model.GetAllTags(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": e.ERROR,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": tags,
		})
	}
}

type CreateNewTagRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateNewTag(c *gin.Context) {
	var req CreateNewTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}

	valid := validation.Validation{}
	valid.Required(req.Name, "name").Message("Must fill Tag Name")
	valid.Required(req.Description, "description").Message("Must fill Tag Description")
	valid.MaxSize(req.Name, 30, "name").Message("Tag Name can not over 30 character")
	valid.MaxSize(req.Description, 300, "description").Message("Tag Description can not over than 300 character")

	if ok, err := model.CheckTagIfExists(req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": e.ERROR,
			"msg":  err.Error(),
		})
		return
	} else if ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.ERROR,
			"msg":  "Tag Name already exists",
		})
		return
	}
	model.CreateNewTag(req.Name, req.Description)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "Create Tag Success",
	})
}

func DeleteTag(c *gin.Context) {
	var id int
	c.ShouldBindUri(&id)
	if _, err := model.DeleteTag(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": e.ERROR,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  "Delete Tag Success",
		})
	}
}
