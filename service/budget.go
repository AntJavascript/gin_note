package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/models"
)

var Budget = new(budgetService)

type budgetService struct{}

// 获取预算
func (c *budgetService) GetBudget(phone string) (dto.Budget, error) {
	detail := dto.Budget{}
	models.DB.Where("account = ?", phone).Find(&detail)
	return detail, nil
}

// 设置预算
func (c *budgetService) SetBudget(amount float64, phone string) error {

	oldVal, _ := Budget.GetBudget(phone)

	detail := dto.Budget{
		Amount: amount,
		Account: phone,
	}

	if oldVal.Amount == 0 {
		// 创建
		if err := models.DB.Create(&detail).Error; err != nil {
			return err
		}
	} else {
		// 修改
		detail.Id = oldVal.Id
		if err := models.DB.Save(&detail).Error; err != nil {
			return err
		}
	}

	return nil
}
