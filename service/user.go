package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/models"
	"example.com/m/v2/tools"
)

var User = new(userService)

type userService struct{}

func (c *userService) Login(params *dto.User) (string, string, error) {
	access_token, err := tools.GenerateToken(params.Phone, params.Password, 2) // 2小时过期
	refresh_token, err := tools.GenerateToken(params.Phone, params.Password, 168) // 7天过期
	return access_token, refresh_token, err
}

// 查询用户是否存在
func (c *userService) QueryUser(params *dto.User) (bool, dto.User) {
	result := dto.User{}
	has := false
	models.DB.Where("phone = ?", params.Phone).Find(&result)
	if result.Phone == params.Phone {
		has = true
	}
	return has, result
}

// 注册新用户
func (c *userService) Register(params *dto.User) error {
  if err := models.DB.Create(&params).Error; err != nil {
		return err
	}
	return nil
}
