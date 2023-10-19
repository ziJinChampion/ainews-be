package main

import (
	"fmt"
	"net/http"

	"github.com/southwind/ainews/adapter"
	"github.com/southwind/ainews/lib"
	"github.com/southwind/ainews/model"
)

func main() {
	model.InitDB(lib.LoadServerConfig())

	router := adapter.InitRouter()

	
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", lib.LoadServerConfig().HTTPPort),
		Handler:        router,
		ReadTimeout:    lib.LoadServerConfig().ReadTimeout,
		WriteTimeout:   lib.LoadServerConfig().WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
