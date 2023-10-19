package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/application"
	"github.com/southwind/ainews/common/code"
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/utils/security"
	"gorm.io/gorm"
)

type Users struct {
	userApplication application.UserAppInterface
}

func NewUsers(userApplication application.UserAppInterface) *Users {
	return &Users{
		userApplication: userApplication,
	}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

func (u *Users) Register(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}
	var user = &entity.User{
		Name:     registerRequest.Name,
		Password: registerRequest.Password,
		Email:    registerRequest.Email,
		Mobile:   registerRequest.Mobile,
	}
	if _, err := user.VaildRegisterInfo(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}
	user.PrepareCreate()
	if _, err := u.userApplication.GetUser(user.Name); err != nil {
		if err == gorm.ErrRecordNotFound {
			_, err := u.userApplication.SaveUser(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": code.ERROR,
					"msg":  err,
				})
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  "User already exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"msg":  "Register Success",
	})

}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *Users) Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}

	var user = &entity.User{
		Name:     loginReq.Name,
		Password: loginReq.Password,
	}

	if _, err := user.ValidLoginInfo(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code.INVALID_PARAMS,
			"msg":  err.Error(),
		})
		return
	}

	res, err := u.userApplication.GetUserByNameAndPassword(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	token, _ := security.GenerateToken(res.Name, res.Role, res.Email, res.Mobile, res.Id)

	c.JSON(http.StatusOK, gin.H{
		"code":    code.SUCCESS,
		"message": "Login Success",
		"data":    token,
	})

}

func FindPassword(c *gin.Context) {
}

func GetUserInfo(c *gin.Context) {

}
