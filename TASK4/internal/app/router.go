package app

import (
	"github.com/gin-gonic/gin"
	"hyperconan.com/blog_sys/internal/app/handlers"
	"hyperconan.com/blog_sys/internal/app/middleware"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	userRouter := Router.Group("/user")
	userRouter.POST("/register", handlers.UserRegister)
	userRouter.POST("/login", handlers.UserLogin)
	userRouter.GET("/test", handlers.Test)

	blogRouter := Router.Group("/blog")
	blogRouter.Use(middleware.JWTAuthMiddleware()) // 应用JWT中间件
	blogRouter.POST("/", handlers.CreatePost)
	blogRouter.PUT("/:post_id", handlers.UpdatePost)
	blogRouter.DELETE("/:post_id", handlers.DeletePost)

	Router.GET("/blog/all", handlers.GetAllPosts)
	Router.GET("/blog/comment/:post_id", handlers.GetCommentsByPostID)

	commentRouter := blogRouter.Group("/comment")
	commentRouter.POST("/", handlers.PostComment)
}
