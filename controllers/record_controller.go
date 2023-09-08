package controllers

import (
	"example.com/m/v2/service"
	"example.com/m/v2/models"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)
var Record = new(recordCtrl)

type recordCtrl struct{}

func (c *recordCtrl) List(ctx *gin.Context) {
	// 参数
	var req *models.Record

	// 调用获取列表方法
	list, count, err := service.Record.GetList(req)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.JsonReturn("","查询失败", 400))
		return
	}
	res := tools.JsonReturn(list, "查询成功", 200)
	res["count"] = count

	// 返回结果
	ctx.JSON(http.StatusOK, res)
}
