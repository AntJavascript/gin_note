package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/models"
	"example.com/m/v2/tools"
	"example.com/m/v2/constants"
)

var User = new(userService)

type userService struct{}

func (c *userService) Login(params *dto.User) (string, string, error) {
	access_token, err := tools.GenerateToken(params.Phone, constants.ACCESSTOKEN)
	refresh_token, err := tools.GenerateToken(params.Phone, constants.REFRESHTOKEN)
	return access_token, refresh_token, err
}

// 查询用户是否存在
func (c *userService) QueryUser(phone string) (bool, dto.User) {
	result := dto.User{}
	has := false
	models.DB.Where("phone = ?", phone).Find(&result)
	if result.Phone == phone {
		has = true
	}
	return has, result
}

// 注册新用户
func (c *userService) Register(params *models.User) error {
  if err := models.DB.Create(&params).Error; err != nil {
		return err
	}
	return nil
}
