package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/m/v2/dto"
	"example.com/m/v2/models"
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
	now := time.Now()
	if date != "" {
		now = tools.StringAsTime(date)
	}
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	date = fmt.Sprintf("%s-%s-%s", year, month, day)

	// 调用获取列表方法
	var req *dto.Record
	list, count, err := service.Record.GetList(req, date)
	if err != nil {
		res = tools.JsonReturn(ctx, "", "查询失败", 400)
	} else {
		res = tools.JsonReturn(ctx, list, "查询成功", 200)
		res["count"] = count
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
			res = tools.JsonReturn(ctx, err, "查询失败", 400)
		} else {
			res = tools.JsonReturn(ctx, detail, "查询成功", 200)
		}
	}
	ctx.JSON(http.StatusOK, res)
}

// 添加记录
func (c *recordCtrl) Add(ctx *gin.Context) {

	res := make(map[string]interface{})

	params := models.Record{}
	err := ctx.BindJSON(&params)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "参数错误", 400)
	} else {
		params.RecordDateUnix = tools.StringAsTampTime(params.RecordDate) // 生成时间戳
		err = service.Record.Add(&params)
		if err != nil {
			res = tools.JsonReturn(ctx, err, "失败", 400)
		} else {
			res = tools.JsonReturn(ctx, "", "成功", 200)
		}
	}

	ctx.JSON(http.StatusOK, res)
}

// 更新单条记录
func (c *recordCtrl) Update(ctx *gin.Context) {

	res := make(map[string]interface{})

	params := models.Record{}
	if ctx.BindJSON(&params) != nil {
		res = tools.JsonReturn(ctx, "", "参数错误", 400)
	} else {
		params.RecordDateUnix = tools.StringAsTampTime(params.RecordDate) // 生成时间戳
		err := service.Record.Update(&params)
		if err != nil {
			res = tools.JsonReturn(ctx, err, "失败", 400)
		} else {
			res = tools.JsonReturn(ctx, "", "成功", 200)
		}
	}

	ctx.JSON(http.StatusOK, res)
}

// 删除单条记录
func (c *recordCtrl) Delete(ctx *gin.Context) {

	res := make(map[string]interface{})

	postId := ctx.Param("id")
	id, _ := strconv.Atoi(postId)

	if postId != "" {
		err := service.Record.Delete(id)

		if err != nil {
			res = tools.JsonReturn(ctx, err, "删除失败", 400)
		} else {
			res = tools.JsonReturn(ctx, "", "删除成功", 200)
		}
	}

	ctx.JSON(http.StatusOK, res)
}
