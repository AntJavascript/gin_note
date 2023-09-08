package controllers

import (
	"github.com/gin-gonic/gin"
)

// 收支记录字段
type Record struct {
	Id         int      // id
	Type       string // income(收入)、expend(支出)
	RecordDate string    // 日期
	Amount     float64   // 金额
	RecordType string    // 收支类型（用于展示图标）
	Remark     string    // 备注
	Account    string    // 账户信息（手机号码）
	Created    time.Time  // 创建时间
}

func (c *Record) List(ctx *gin.Context) {
	// 参数
	var req *dto.UserPageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用获取列表方法
	list, count, err := service.Record.GetList(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "查询成功",
		Count: count,
	})
}
