package controllers
 
import (
	"fmt"
	"net/http"
	"time"
	"strconv"

	"example.com/m/v2/service"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
)

var TotalStruct = new(totalCtrl)

type totalCtrl struct {
}

// 获取某日数据
func (c *totalCtrl) GetDay(ctx *gin.Context) {
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
	incomeCount, expendCount, err := service.Total.GetDayTotal(tools.StringAsTime(date).Unix())
	if err != nil {
		res = tools.JsonReturn(ctx, err, "失败", 400)
	} else {
		res = tools.JsonReturn(ctx, "", "成功", 200)
		res["incomeCount"] = incomeCount
		res["expendCount"] = expendCount
	}
	ctx.JSON(http.StatusOK, res)
	return
}

// 获取某月数据
func (c *totalCtrl) GetMonth(ctx *gin.Context) {
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

	fmt.Println(maxDay)

	incomeCount, expendCount, err := service.Total.GetMonthOrYearTotal(startDate, endDate)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "失败", 400)
	} else {
		res = tools.JsonReturn(ctx, "", "成功", 200)
		res["incomeCount"] = incomeCount // 收入
		res["expendCount"] = expendCount // 支出
	}
	ctx.JSON(http.StatusOK, res)
	return
}

// 获取某年数据
func (c *totalCtrl) GetYear(ctx *gin.Context) {
	res := make(map[string]interface{})

	year := ctx.Param("year")

	var startDate int64 // 本年第一天
	var endDate int64   // 本年最后一天

	if year == "" {
		// 默认本年
		now := time.Now()
		year = now.Format("2006")
	}

	startDate = tools.StringAsTime(year + "-01-01").Unix()
	endDate = tools.StringAsTime(year + "-12-31").Unix()

	incomeCount, expendCount, err := service.Total.GetMonthOrYearTotal(startDate, endDate)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "失败", 400)
	} else {
		res = tools.JsonReturn(ctx, "", "成功", 200)
		res["incomeCount"] = incomeCount // 收入
		res["expendCount"] = expendCount // 支出
	}

	ctx.JSON(http.StatusOK, res)
	return
}
