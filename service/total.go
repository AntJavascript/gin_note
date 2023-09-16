package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/models"
)

var Total = new(totalService)

type totalService struct{}

// 统计当天支出和收入
func (c *totalService) GetDayTotal(date int64) (float64, float64, error) {
	
	var result []dto.Record

	var incomeCount float64 // 收入
	var expendCount float64 // 支出

	models.DB.Where("record_date = ?", date).Find(&result)

	for i := 0; i < len(result) - 1; i++ {
		if result[i].Type == "income" {
			incomeCount += result[i].Amount
		} else {
			expendCount += result[i].Amount
		}
	}
	
	return incomeCount, expendCount, nil
}

// 统计当天支出和收入
func (c *totalService) GetMonthTotal(startDate, endDate int64) (float64, float64, error) {
	
	var result []dto.Record

	var incomeCount float64 // 收入
	var expendCount float64 // 支出

	models.DB.Where("record_date_unix >= ? AND record_date_unix <= ?", startDate, endDate).Find(&result)

	for i := 0; i < len(result) - 1; i++ {
		if result[i].Type == "income" {
			incomeCount += result[i].Amount
		} else {
			expendCount += result[i].Amount
		}
	}

	return incomeCount, expendCount, nil
}
