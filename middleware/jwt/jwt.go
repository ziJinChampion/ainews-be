package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/pkg/e"
	"github.com/southwind/ainews/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Request.Header("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if t, ok := claims.ExpiresAt.(*time.Time) && t < time.Now()  {
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
