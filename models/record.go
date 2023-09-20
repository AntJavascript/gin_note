package models

import (
	"gorm.io/gorm"
)

// 收支记录字段
type Record struct {
	gorm.Model
	Type       string  `gorm:"not null"`    // income(收入)、expend(支出)
	RecordDate string  `gorm:"not null"`    // 日期
	RecordDateUnix int64  `gorm:"not null"`    // 时间戳
	Amount     float64  `gorm:"not null"`   // 金额
	RecordType string  `gorm:"not null"`    // 收支类型（用于展示图标）
	Remark     string    // 备注
	Account    string  `gorm:"type:varchar(11);not null"`    // 账户信息（手机号码）
}
