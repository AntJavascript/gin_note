package controllers

import (
	"strconv"
	"example.com/m/v2/service"
	"example.com/m/v2/dto"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)
var Record = new(recordCtrl)

type recordCtrl struct{}

func (c *recordCtrl) List(ctx *gin.Context) {

	res := make(map[string]interface{})

	defer ctx.JSON(http.StatusOK, res)

	var req *dto.Record

	// 调用获取列表方法
	list, count, err := service.Record.GetList(req)
	if err != nil {
		return
	}
	res = tools.JsonReturn(list, "查询成功", 200)
	token, err := tools.GenerateToken("15817351609", "shi@465608")
	res["count"] = count
	res["token"] = token

}

func (c *recordCtrl) Detail(ctx *gin.Context) {

	postId := ctx.Param("id")
	id, _ := strconv.Atoi(postId)
	var req *dto.Record
	
	res := make(map[string]interface{})

	defer ctx.JSON(http.StatusOK, res)

	if postId != "" {
		detail, err := service.Record.Detail(req, id)
		if err != nil {
			res = tools.JsonReturn("", "查询失败", 400)
		} else {
			res = tools.JsonReturn(detail, "查询成功", 200)
		}
	}
}
