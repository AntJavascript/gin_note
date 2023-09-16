package models

import (
	"gorm.io/gorm"
)

// 用户表字段
type User struct {
	gorm.Model
	UserName string
	Password string
	Email    string
	Phone    string
	Face     string // 用户头像
	Status   int    // 账号状态
}
