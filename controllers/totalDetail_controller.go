package controllers

import (
	"net/http"
	"strconv"
	"time"

	"example.com/m/v2/service"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
)

var TotalDetailStruct = new(totalDetailCtrl)

type totalDetailCtrl struct {
}

// 获取某月数据统计详情
func (c *totalDetailCtrl) GeTotalMonthDetail(ctx *gin.Context) {
	res := make(map[string]interface{})

	year := ctx.Query("year")   // 年
	month := ctx.Query("month") // 月

	var startDate int64 // 月初第一天
	var endDate int64   // 月底最后一天

	if year == "" || month == "" {
		// 默认本年、本月
		now := time.Now()
		year = now.Format("2006")
		month = now.Format("01")
	}

	startDate = tools.StringAsTime(year + "-" + month + "-" + "01").Unix()
	maxDay := strconv.Itoa(tools.GetMonthDaxDay(year, month)) // 月最后一天
	endDate = tools.StringAsTime(year + "-" + month + "-" + maxDay).Unix()

	list, err := service.Total.GetMonthDetail(startDate, endDate)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "失败", 400)
	} else {
		res = tools.JsonReturn(ctx, "", "成功", 200)
		res["data"] = list
	}
	ctx.JSON(http.StatusOK, res)
	return
}
