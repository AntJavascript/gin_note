package models

import (
	"gorm.io/gorm"
)

// 预算字段
type Budget struct {
	gorm.Model
	Account    string  `gorm:"not null"`    // 账户信息（手机号码）
	Amount     float64 string  `gorm:"default:0;not null"`   // 金额
}
