package main

import (
	"fmt"
	"net/http"

	"github.com/fy026/weather/gateway/pkg/setting"
	"github.com/fy026/weather/gateway/routers"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
