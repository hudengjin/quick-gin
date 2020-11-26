package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/controllers"
)
// Api 定义 api 路由注册
func Api(router *gin.Engine) {
	apiGroup := router.Group("api")
	{
		apiGroup.GET("/test", controllers.ApiIndex)
	}

	api := router.Group("/api")
	api.GET("/index", controllers.ApiIndex)

	// 中间件
	// api.Use

}