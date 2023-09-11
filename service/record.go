package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/models"
)

var Record = new(recordService)

type recordService struct{}

func (c *recordService) GetList(req *dto.Record, date string) ([]dto.Record, int, error) {
	// 数据处理
	var result []dto.Record
	models.DB.Where("created = ?", date).Find(&result)
	return result, len(result), nil
}

func (c *recordService) Detail(req *dto.Record, id int) (dto.Record, error) {
	// 数据处理
	detail := dto.Record{
		Id: id,
	}
	models.DB.Find(&detail)
	return detail, nil
}

func (c *recordService) Add(params *dto.Record) error {

	if err := models.DB.Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (c *recordService) Update(params *dto.Record) error {

	if err := models.DB.Save(&params).Error; err != nil {
		return err
	}
	return nil
}

func (c *recordService) Delete(id int) error {
	detail := dto.Record{
		Id: id,
	}
	if err := models.DB.Delete(&detail).Error; err != nil {
		return err
	}
	return nil
}
