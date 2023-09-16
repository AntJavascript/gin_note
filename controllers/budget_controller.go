package controllers

import (
	"net/http"

	"example.com/m/v2/service"
	"example.com/m/v2/tools"
	"example.com/m/v2/dto"
	"github.com/gin-gonic/gin"
)

var BudgetStruct = new(budgetCtrl)

type budgetCtrl struct {
}

// 获取预算
func (c *budgetCtrl) GetBudget(ctx *gin.Context) {
	res := make(map[string]interface{})

	access_token, _ :=  tools.GetHearToken(ctx)
	phone := tools.GetTokenAccount(access_token)

	if phone == "" {
		res = tools.JsonReturn(ctx, "", "失败", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}
	data, err := service.Budget.GetBudget(phone)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "失败", 400)
	} else {
		res = tools.JsonReturn(ctx, data, "成功", 200)
	}
	ctx.JSON(http.StatusOK, res)
	return
}

//设置预算
func (c *budgetCtrl) SetBudget(ctx *gin.Context) {
	res := make(map[string]interface{})

	access_token, _ :=  tools.GetHearToken(ctx)
	phone := tools.GetTokenAccount(access_token)

	// 解析前端body参数
	params := dto.Budget{}
	ctx.BindJSON(&params)

	if phone == "" {
		res = tools.JsonReturn(ctx, "", "失败", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}
	err := service.Budget.SetBudget(params.Amount, phone)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "失败", 400)
	} else {
		res = tools.JsonReturn(ctx, "", "成功", 200)
	}
	ctx.JSON(http.StatusOK, res)
	return
	return
}
