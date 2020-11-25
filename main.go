package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/config"
	"github.com/huprince/quick-gin/modules/server"
	"github.com/huprince/quick-gin/routes"
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := routes.InitRouter()
	server.Run(router)

}