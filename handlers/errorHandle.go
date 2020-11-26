package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/huprince/quick-gin/modules/log"
	"github.com/huprince/quick-gin/util"
)

// ErrorHandle 错误处理
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func ()  {
			if err := recover(); err != nil {
				logger.Logger.Error(util.InterfaceToString(err))
				if errMsg, ok := err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": errMsg,
					})
					return
				} 
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg": "System error!",
				})
				return
			}
		}()
		c.Next()
	}
}