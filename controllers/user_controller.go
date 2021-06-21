package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/connections"
	"github.com/huprince/quick-gin/models"
	"github.com/huprince/quick-gin/modules/auth"
	logger "github.com/huprince/quick-gin/modules/log"
	"github.com/huprince/quick-gin/services"
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
		Password: services.BcryptPassword(password),
	}
	db.Save(user)
	
	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"auth_type": "basic",
		"token": auth.GenerateBasicToken(email, password),
		"msg": "created successful!",
	})

}

// Login 登录接口
func Login(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg": err.Error(),
		})
		return
	}

	email := c.PostForm("email")
	password := c.PostForm("password")

	r := db.Where(&models.User{Email: email,}).First(user)

	if r.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg": "Not found current user!",
		})
		return
	}
	
	if !services.IsCorrectPassword(user.Password, password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400, 
			"msg": "Password is wrong!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "Login successful!",
		"auth_type": "basic",
		"token": auth.GenerateBasicToken(email, password),
	})

}