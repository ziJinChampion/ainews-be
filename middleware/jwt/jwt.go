package jwt

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/pkg/constant"
	"github.com/southwind/ainews/pkg/e"
	"github.com/southwind/ainews/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := utils.ParseToken(tokenStr)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				log.Fatal(err)
			} else if claims.ExpiresAt.Before(time.Now()) {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := utils.ParseToken(tokenStr)
			log.Default().Print(claims)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				log.Fatal(err)
			} else if claims.Role != constant.ADMIN_USER {
				code = e.ERROR_AUTHENTICATION
			} else if claims.ExpiresAt.Before(time.Now()) {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
