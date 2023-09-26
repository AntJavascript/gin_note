package models

import (
	"gorm.io/gorm"
)

// 收支记录字段 gorm: 数据库相关、json:前端传参相关
type Record struct {
	gorm.Model
	Type           string  `gorm:"default:'expend';not null" json:"type"`    // income(收入)、expend(支出)
	RecordDate     string  `gorm:"not null" json:"record_date"`              // 日期
	RecordDateUnix int64   `gorm:"not null" json:"record_date_unix"`         // 时间戳
	Amount         float64 `gorm:"not null" json:"amount"`                   // 金额
	RecordType     string  `gorm:"not null" json:"record_type"`              // 收支类型（用于展示图标）
	RecordTypeName string  `gorm:"not null" json:"record_type_name"`         // 收支类型名称
	Remark         string  `json:"remark"`                                   // 备注
	Account        string  `gorm:"type:varchar(11);not null" json:"account"` // 账户信息（手机号码）
}
