package models

import (
	"time"
)

// 用户表字段
type User struct {
	Id       int
	UserName string
	Password string
	Email    string
	Created  time.Time `gorm:"autoUpdateTime:false"` // 注册时间
	Phone    string
	Face     string // 用户头像
	Status   int    // 账号状态
}
