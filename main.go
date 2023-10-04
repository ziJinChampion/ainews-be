package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/lib"
)

func main() {

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", lib.LoadServerConfig().HTTPPort),
		Handler:        router,
		ReadTimeout:    lib.LoadServerConfig().ReadTimeout,
		WriteTimeout:   lib.LoadServerConfig().WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
