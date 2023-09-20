package models

import (
	"gorm.io/gorm"
)

// 用户表字段
type User struct {
	gorm.Model
	UserName string
	Password string `gorm:"size:255;not null"`
	Email    string
	Phone    string `gorm:"type:varchar(11);not null"`
	Face     string // 用户头像
	Status   int    // 账号状态
}
