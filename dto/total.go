package dto

// 统计数据返回字段
type Total struct {
	Date 					int `json:"month"`
	Amount          float64   `json:"amount"`
	Type           string  `json:"type"` // income(收入)、expend(支出)
}
