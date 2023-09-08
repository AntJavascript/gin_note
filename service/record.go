package service
import (
	"example.com/m/v2/models"
)

var Record = new(recordService)

type recordService struct{}

func(c *recordService) GetList(req *models.Record) ([]models.Record, int, error) {
	// 数据处理
	var result []models.Record
	// 总数
	var count int = 10
	for i := 0; i <= count; i++ {
		item := models.Record{
			Id: i,
			Type: "income",
		}
		// 加入数组
		result = append(result, item)
	}

	return result, count, nil
}
