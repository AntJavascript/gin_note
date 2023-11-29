package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/m/v2/service"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
)

// 获取某月数据统计详情
func (c *totalCtrl) GeTotalMonthDetail(ctx *gin.Context) {
  res := make(map[string]interface{})

	year := ctx.Query("year")   // 年
	month := ctx.Query("month") // 月
}
