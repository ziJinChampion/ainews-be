package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/southwind/ainews/lib"
	"github.com/southwind/ainews/middleware/jwt"
	v1 "github.com/southwind/ainews/routers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(lib.LoadServerConfig().RunMode)

	userApi := r.Group("/")
	{
		userApi.POST("/login", v1.Login)
		userApi.POST("/register", v1.Register)
		userApi.POST("/find-password", v1.FindPassword)
	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/info", v1.GetUserInfo)
		apiv1.GET("/tags", v1.GetAllTags)
		apiv1.POST("/tags", v1.CreateNewTag)
	}

	adminApi := r.Group("/api/admin")
	adminApi.Use(jwt.AdminJWT())
	{
		adminApi.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
