package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/controllers"
)
// API 定义 api 路由注册
func API(router *gin.Engine) {
	apiGroup := router.Group("api")
	{
		apiGroup.GET("/test", controllers.APIIndex)
		apiGroup.POST("/users", controllers.Register)
		apiGroup.POST("/login", controllers.Login)
	}

	api := router.Group("/api")
	api.GET("/index", controllers.APIIndex)

	// 中间件
	// api.Use

}