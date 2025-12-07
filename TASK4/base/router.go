package base

import "github.com/gin-gonic/gin"
import "hyperconan.com/blog_sys/modules"

var Router *gin.Engine

func init() {
	Router = gin.Default()

	userRouter := Router.Group("/user")
	userRouter.POST("/register", modules.UserRegister)
	userRouter.POST("/login", modules.UserLogin)
	userRouter.GET("/test", modules.Test)
}
