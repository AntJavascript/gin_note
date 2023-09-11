package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/m/v2/dto"
	"example.com/m/v2/service"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
)

var Record = new(recordCtrl)

type recordCtrl struct{}

// 根据日期查询列表
func (c *recordCtrl) List(ctx *gin.Context) {
	res := make(map[string]interface{})
	
	date := ctx.Query("date") // 请求日期
	if date == "" {
		now := time.Now()
		year := now.Format("2006")
		month := now.Format("01")
		day := now.Format("02")
		date = fmt.Sprintf("%s-%s-%s", year, month, day)
	}
	fmt.Println(date)

	// 调用获取列表方法
	var req *dto.Record
	list, count, err := service.Record.GetList(req, date)
	if err != nil {
		res = tools.JsonReturn("", "查询失败", 400)
	} else {
		res = tools.JsonReturn(list, "查询成功", 200)
		token, err := tools.GenerateToken("15817351609", "shi@465608")
		res["count"] = count
		res["token"] = token
	}
	
	ctx.JSON(http.StatusOK, res)
}

// 查询单条记录详情
func (c *recordCtrl) Detail(ctx *gin.Context) {
	postId := ctx.Param("id")
	id, _ := strconv.Atoi(postId)
	res := make(map[string]interface{})

	if postId != "" {
		var req *dto.Record
		detail, err := service.Record.Detail(req, id)
		if err != nil {
			res = tools.JsonReturn("", "查询失败", 400)
		} else {
			res = tools.JsonReturn(detail, "查询成功", 200)
		}
	}
	ctx.JSON(http.StatusOK, res)
}
