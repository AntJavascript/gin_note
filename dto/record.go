package dto

import (
	"time"
)

// 收支记录字段
type Record struct {
	Id         int      `json:"id"`      // id
	Type       string   `json:"type"` // income(收入)、expend(支出)
	RecordDate string   `json:"record_date"`    // 日期
	Amount     float64  `json:"amount"`   // 金额
	RecordType string   `json:"record_type"`    // 收支类型（用于展示图标）
	Remark     string   `json:"remark"`    // 备注
	Account    string   `json:"account"`    // 账户信息（手机号码）
	Created    time.Time  `json:"created"`  // 创建时间
}
