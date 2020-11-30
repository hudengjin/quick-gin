package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/connections"
	"github.com/huprince/quick-gin/models"
	logger "github.com/huprince/quick-gin/modules/log"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	db, err = connections.InitDB()
	if err != nil {
		logger.Logger.Error("Database init connection fail!")
		log.Fatalln(err)
	}
}

// Register 用户注册
func Register(c *gin.Context) {
	username := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := &models.User{
		Name: username,
		Email: email,
		Password: password,
	}
	db.Save(user)
	
	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg": "created successful!",
	})

}