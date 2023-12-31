package adapter

import (
	"github.com/gin-gonic/gin"

	"github.com/southwind/ainews/adapter/middleware"
	v1 "github.com/southwind/ainews/adapter/v1"
	lib "github.com/southwind/ainews/common/config"
	"github.com/southwind/ainews/infrastructure/persistence"
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
	tags := v1.NewTag(daos.Tag)
	articles := v1.NewArticles(daos.Article, daos.Tag, daos.User)

	{
		indexApi.POST("/login", users.Login)
		indexApi.POST("/register", users.Register)
		indexApi.GET("/tags", tags.GetTags)
		indexApi.GET("/articles", articles.GetArticles)
	}
	authApi := r.Group("/api")
	authApi.Use(middleware.JWT())
	{
		authApi.POST("/tags", tags.CreateTag)
		authApi.POST("/articles", articles.CreateArticle)
	}

	return r
}
