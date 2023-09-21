package dto

// 登录\注册
type User struct {
	Phone       string   `json:"phone"`
	Password   string   `json:"password"`
	Email   string   `json:"email"`
}
