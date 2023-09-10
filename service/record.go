package service

import (
	"fmt"

	"example.com/m/v2/dto"
	"example.com/m/v2/models"
)

var Record = new(recordService)

type recordService struct{}

func (c *recordService) GetList(req *dto.Record) ([]dto.Record, int, error) {
	// 数据处理
	var result []dto.Record
	// 总数
	var count int = 10
	for i := 0; i <= count; i++ {
		item := dto.Record{
			Id:   i,
			Type: "income",
		}
		// 加入数组
		result = append(result, item)
	}

	return result, count, nil
}

func (c *recordService) Detail(req *dto.Record, id int) (dto.Record, error) {
	// 数据处理
	detail := dto.Record{
		Id: id,
	}
	fmt.Println("models.DBmodels.DBmodels.DBmodels.DB", models.DB)
	fmt.Println("models.DBmodels.DBmodels.DBmodels.DB", models.DB)
	models.DB.Find(&detail)
	return detail, nil
}
