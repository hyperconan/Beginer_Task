package main

import (
	"fmt"
	"net/http"
	"time"

	"hyperconan.com/blog_sys/internal/app"
	"hyperconan.com/blog_sys/internal/pkg/logger"
)

func main() {
	logger.Init()
	defer logger.Sync()

	fmt.Println("hello world!")
	logger.S.Infow("starting server", "addr", ":7913")
	//启动方式1
	//app.Router.Run(":7913")

	//启动方式2
	s := &http.Server{
		Addr:         ":7913",
		Handler:      app.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		logger.S.Errorw("server stopped with error", "err", err)
	}

	// 启动方式3
	//http.ListenAndServe(":7913", app.Router)
}
