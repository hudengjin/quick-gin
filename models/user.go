package models

import "gorm.io/gorm"

// User 模型定义
type User struct {
	gorm.Model
	Name string `gorm:"size:128"`
	Email string `gorm:"unique;size:256" form:"email" binding:"required"`
	Password string `gorm:"size:128" form:"password" binding:"required"`
	ResetToken string `gorm:"size:128"`
}