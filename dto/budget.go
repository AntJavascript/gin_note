package dto

// 预算
type Budget struct {
	Id 					int `json:"id"`
	Amount       float64   `json:"amount"`
	Account			string   `json:"account"`
}
