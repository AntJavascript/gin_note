package models

import (
	"gorm.io/gorm"
)

// 收支记录字段
type Record struct {
	gorm.Model
	Type       string    // income(收入)、expend(支出)
	RecordDate string    // 日期
	RecordDateUnix int64    // 时间戳
	Amount     float64   // 金额
	RecordType string    // 收支类型（用于展示图标）
	Remark     string    // 备注
	Account    string    // 账户信息（手机号码）
}
