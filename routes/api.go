package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/controllers"
)

func Api(router *gin.Engine)  {
	apiGroup := router.Group("api")
	{
		apiGroup.GET("/test", controllers.ApiIndex)
	}

	api := router.Group("/api")
	api.GET("/index", controllers.ApiIndex)

	// 中间件
	// api.Use

}