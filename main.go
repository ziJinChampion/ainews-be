package main

import (
	"fmt"
	"net/http"

	"github.com/southwind/ainews/lib"
	"github.com/southwind/ainews/routers"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", lib.LoadServerConfig().HTTPPort),
		Handler:        router,
		ReadTimeout:    lib.LoadServerConfig().ReadTimeout,
		WriteTimeout:   lib.LoadServerConfig().WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
