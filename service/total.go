package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/models"
	"example.com/m/v2/tools"
)

var Total = new(totalService)

type totalService struct{}

// 统计当天支出和收入
func (c *totalService) GetDayTotal(date int64) (float64, float64, error) {

	var result []dto.Record

	var incomeCount float64 // 收入
	var expendCount float64 // 支出

	models.DB.Where("record_date_unix = ?", date).Find(&result)

	for i := 0; i < len(result); i++ {
		if result[i].Type == "income" {
			incomeCount += result[i].Amount
		} else {
			expendCount += result[i].Amount
		}
	}
	return tools.Decimal(incomeCount, 2), tools.Decimal(expendCount, 2), nil
}

// 统计某月或者某年的记录
func (c *totalService) GetMonthOrYearTotal(startDate, endDate int64) (float64, float64, error) {
	var result []dto.Record

	var incomeCount float64 // 收入
	var expendCount float64 // 支出

	models.DB.Where("record_date_unix BETWEEN ? AND ?", startDate, endDate).Find(&result)

	for i := 0; i < len(result); i++ {
		if result[i].Type == "income" {
			incomeCount += result[i].Amount
		} else {
			expendCount += result[i].Amount
		}
	}
	return tools.Decimal(incomeCount, 2), tools.Decimal(expendCount, 2), nil
}

// 统计某月数据明细
func (c *totalService) GetMonthDetail(startDate, endDate int64) ([]dto.Total, error) {
	var result []dto.Total
	models.DB.Row("select day(FROM_UNIXTIME(record_date_unix)) as day,sum(amount)as amount from record where record_date_unix >= startDate and record_date_unix <= startDate group by day").Scan(&result)
	return result, nil
}

