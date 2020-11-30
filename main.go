package main

import (
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/config"
	"github.com/huprince/quick-gin/connections"
	"github.com/huprince/quick-gin/models"
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

	if config.GetEnv().DbConfig.AutoMigrate {
		db, err := connections.InitDB()
		if err != nil {
			log.Fatalln(err)
		}
		db.AutoMigrate(&models.User{})
	}

	router := routes.InitRouter()
	server.Run(router)

}