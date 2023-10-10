package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/southwind/ainews/lib"
	v1 "github.com/southwind/ainews/routers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(lib.LoadServerConfig().RunMode)

	apiv1 := r.Group("/api/v1")
	{

		apiv1.POST("/login", v1.Login)
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/find-password", v1.FindPassword)
		apiv1.GET("/info", v1.GetUserInfo)

	}

	return r
}
