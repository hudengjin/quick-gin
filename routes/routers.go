package routes

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/config"
)

func InitRouter() *gin.Engine{
	router := gin.New()

	if config.GetEnv().Debug {
		pprof.Register(router)
	}

	// 中间件
	router.Use(gin.Logger())
	// router.Use()

	// 全局路由异常信息处理
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "404",
			"msg": "Request route not found!",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": "400",
			"msg": "Request Method is not incorrect!",
		})
	})

	// 自定义路由
	Api(router)

	return router
}