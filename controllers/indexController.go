package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg": "Hello Quick gin!",
	})
}