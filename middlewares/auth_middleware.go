package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/modules/auth"
	"github.com/huprince/quick-gin/modules/log"
)

func BasicAuth() gin.HandlerFunc {
	return func (c *gin.Context)  {
		if header := c.Request.Header.Get("Authorization"); header != "" {
			if idx := strings.Index(header, " "); strings.ToLower(header[:idx]) == "basic" {
				if ok, err := auth.CheckBaiscAuth(strings.Trim(header[idx:], " ")); ok {
					c.Next()
				} else {
					log.InitLogger().Error(err.Error())
				}
			}
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
}