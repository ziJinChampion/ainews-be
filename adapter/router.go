package adapter

import (
	"github.com/gin-gonic/gin"

	"github.com/southwind/ainews/adapter/middleware"
	v1 "github.com/southwind/ainews/adapter/v1"
	"github.com/southwind/ainews/infrastructure/persistence"
	"github.com/southwind/ainews/lib"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//Cors
	r.Use(middleware.CORSMiddleware())

	gin.SetMode(lib.LoadServerConfig().RunMode)

	indexApi := r.Group("/")

	daos, err := persistence.NewDAO(lib.LoadServerConfig())
	if err != nil {
		panic(err)
	}
	daos.Migrate()

	users := v1.NewUsers(daos.User)

	{
		indexApi.POST("/login", users.Login)
		indexApi.POST("/register", users.Register)
	}
	// apiv1WithoutAuth := r.Group("/api/v1")
	// {
	// 	apiv1WithoutAuth.GET("/tags", v1.GetAllTags)
	// 	apiv1WithoutAuth.GET("/articles", v1.GetAllArticles)
	// }
	// apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	// {
	// 	apiv1.GET("/info", v1.GetUserInfo)
	// 	apiv1.POST("/tags", v1.CreateNewTag)
	// 	apiv1.POST("/articles", v1.CreateNewArticle)
	// }

	// adminApi := r.Group("/api/admin")
	// adminApi.Use(jwt.AdminJWT())
	// {
	// 	adminApi.DELETE("/tags/:id", v1.DeleteTag)
	// }

	return r
}
