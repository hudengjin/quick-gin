package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/modules/log"
)

// ApiIndex index 路由控制器
func ApiIndex(c *gin.Context) {
	log.Logger.Info("Quick gin log info")
	log.Logger.Warn("Quick gin log warn")
	log.Logger.Error("Quick gin log error")
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg": "Hello Quick gin!",
	})
}

