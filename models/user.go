package models

import "gorm.io/gorm"

// User 模型定义
type User struct {
	gorm.Model
	Name string `gorm:"size:128"`
	Email string `gorm:"size:256"`
	Password string `gorm:"size:128"`
	ResetToken string `gorm:"size:128"`
}