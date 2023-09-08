package models

import (
	"time"
)

// 收支记录字段
type Record struct {
	Id         int      `gorm:"primaryKey"`      // id
	Type       string    // income(收入)、expend(支出)
	RecordDate string       // 日期
	Amount     float64     // 金额
	RecordType string       // 收支类型（用于展示图标）
	Remark     string       // 备注
	Account    string       // 账户信息（手机号码）
	Created    time.Time    // 创建时间
}

func (table *Record) TableName() string {
	return "Record"
}
