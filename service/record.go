package service
import (
	"example.com/m/v2/dto"
)

var Record = new(recordService)

type recordService struct{}

func(c *recordService) GetList(req *dto.Record) ([]dto.Record, int, error) {
	// 数据处理
	var result []dto.Record
	// 总数
	var count int = 10
	for i := 0; i <= count; i++ {
		item := dto.Record{
			Id: i,
			Type: "income",
		}
		// 加入数组
		result = append(result, item)
	}

	return result, count, nil
}
