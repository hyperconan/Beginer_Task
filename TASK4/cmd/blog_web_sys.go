package main

import (
	"fmt"
	"net/http"
	"time"

	"hyperconan.com/blog_sys/internal/app"
)

func main() {
	fmt.Println("hello world!")
	//启动方式1
	//app.Router.Run(":7913")

	//启动方式2
	s := &http.Server{
		Addr:         ":7913",
		Handler:      app.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()

	// 启动方式3
	//http.ListenAndServe(":7913", app.Router)
}
