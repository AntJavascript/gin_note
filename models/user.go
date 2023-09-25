package models

import (
	"gorm.io/gorm"
)

// 用户表字段
type User struct {
	gorm.Model
	UserName string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"varchar(100);not null"`
	Phone    string  `gorm:"varchar(11);not null;unique"`
	Face     string // 用户头像
	Status   int    // 账号状态
}
