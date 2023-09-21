package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/models"
)

var Record = new(recordService)

type recordService struct{}

// 根据日期获取列表
func (c *recordService) GetList(req *dto.Record, date string) ([]dto.Record, int, error) {
	var result []dto.Record
	models.DB.Where("record_date = ?", date).Find(&result)
	return result, len(result), nil
}

// 根据id获取详情
func (c *recordService) Detail(req *dto.Record, id int) (dto.Record, error) {
	detail := dto.Record{}
	models.DB.Find(&detail, id)
	return detail, nil
}

// 添加数据
func (c *recordService) Add(params *dto.Record) error {
	err := models.DB.Create(&params).Error
	if err != nil {
		return err
	}
	return nil
}

// 更新数据
func (c *recordService) Update(params *dto.Record) error {
	err := models.DB.Save(&params).Error
	// err := models.DB.Updates(&params).Error
	if  err != nil {
		return err
	}
	return nil
}

// 根据id删除数据
func (c *recordService) Delete(id int) error {
	detail := dto.Record{
		Id: id,
	}
	if err := models.DB.Delete(&detail).Error; err != nil {
		return err
	}
	return nil
}
