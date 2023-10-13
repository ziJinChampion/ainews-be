package v1

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/model"
	"github.com/southwind/ainews/pkg/e"
	"github.com/southwind/ainews/pkg/utils"
)

func Login(c *gin.Context) {
	data, _ := c.GetRawData()
	var body map[string]string
	json.Unmarshal(data, &body)

	name := body["name"]
	password := body["password"]

	valid := validation.Validation{}
	valid.Required(name, "name").Message("Must fill UserName")
	valid.Required(password, "password").Message("Must fill Password")
	valid.MaxSize(name, 30, "name").Message("UserName can not over 30 character	")
	valid.MaxSize(password, 30, "password").Message("Password can not over than 30 character ")

	code := e.INVALID_PARAMS
	res := make(map[string]interface{})
	message := "Login Success"
	if !valid.HasErrors() {
		if ok, err := model.ValidUserInfo(name, password); ok && err == nil {
			code = e.SUCCESS
			tokenStr, err := utils.GenerateToken(name, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				res["token"] = tokenStr
			}
		} else if err != nil {
			code = e.ERROR
			message = err.Error()
		} else {
			code = e.ERROR_AUTH
			message = "UserName or Password is wrong"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    res,
	})

}

func Register(c *gin.Context) {

	data, _ := c.GetRawData()

	var body map[string]string
	json.Unmarshal(data, &body)

	name := body["name"]
	password := body["password"]
	mobile := body["mobile"]
	email := body["email"]

	valid := validation.Validation{}
	valid.Required(name, "name").Message("Must fill UserName")
	valid.Required(password, "password").Message("Must fill Password")
	valid.MaxSize(name, 30, "name").Message("UserName can not over 30 character	")
	valid.MaxSize(password, 30, "password").Message("Password can not over than 30 character ")
	valid.MinSize(name, 8, "name").Message("UserName mush over than 8 character	")
	valid.MinSize(password, 8, "password").Message("Password must over than 8 character	")

	code := e.INVALID_PARAMS

	if valid.HasErrors() {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": valid.Errors[0].String(),
		})
		return
	}

	if u, err := model.GetUserInfo(map[string]string{"user_name": name}); err == nil && u.Id > 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    e.INVALID_PARAMS,
			"message": "this user name already being used",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    e.ERROR,
			"message": err.Error(),
		})
		return
	}

	if _, err := model.RegisterUser(name, password, mobile, email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    e.ERROR,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "Register Success",
	})
}

func FindPassword(c *gin.Context) {
}

func GetUserInfo(c *gin.Context) {

}
