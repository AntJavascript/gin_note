package controllers

import (
	"example.com/m/v2/service"
	"example.com/m/v2/dto"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)
var Record = new(recordCtrl)

type recordCtrl struct{}

func (c *recordCtrl) List(ctx *gin.Context) {
	// 参数
	var req *dto.Record

	// 调用获取列表方法
	list, count, err := service.Record.GetList(req)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.JsonReturn("","查询失败", 400))
		return
	}
	res := tools.JsonReturn(list, "查询成功", 200)
	token, err := tools.GenerateToken("15817351609", "shi@465608")
	res["count"] = count
	res["token"] = token

	// 返回结果
	ctx.JSON(http.StatusOK, res)
}
