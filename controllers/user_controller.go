package controllers

import (
	"net/http"

	"example.com/m/v2/dto"
	"example.com/m/v2/service"
	"example.com/m/v2/tools"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

var User = new(userCtrl)

type userCtrl struct{}

// 登录
func (c *userCtrl) Login(ctx *gin.Context) {
	res := make(map[string]interface{})

	// 解析前端body参数
	params := dto.User{}
	ctx.BindJSON(&params)

	if !verification(&params) {
		res = tools.JsonReturn(ctx, "", "输入不合法", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}

	// 查询用户是否存在
	hasUser, userInfo := service.User.QueryUser(&params)

	if !hasUser {
		res = tools.JsonReturn(ctx, "", "用户不存在", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}

	if userInfo.Password != params.Password {
		res = tools.JsonReturn(ctx, "", "密码不正确", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}

	access_token, refresh_token, err := service.User.Login(&params)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "登录失败", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = tools.JsonReturn(ctx, "", "登录成功", 200)
	res["access_token"] = access_token
	res["refresh_token"] = refresh_token
	ctx.JSON(http.StatusOK, res)
}

// 注册
func (c *userCtrl) Register(ctx *gin.Context) {
	res := make(map[string]interface{})

	// 解析前端body参数
	params := models.User{}
	ctx.BindJSON(&params)

	if !verification(&params) {
		res = tools.JsonReturn(ctx, "", "输入不合法", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}

	// 查询用户是否存在
	hasUser, _ := service.User.QueryUser(params.Phone)

	if hasUser {
		res = tools.JsonReturn(ctx, "", "手机号已注册", 400)
		ctx.JSON(http.StatusOK, res)
		return
	}

	err := service.User.Register(&params)
	if err != nil {
		res = tools.JsonReturn(ctx, err, "注册失败", 400)
	} else {
		res = tools.JsonReturn(ctx, "", "注册成功", 200)
	}
	ctx.JSON(http.StatusOK, res)
}

// 输入非空验证
func verification(params *models.User) bool {
	pass := true
	if params.Phone == "" || params.Password == "" {
		pass = false
	}
	return pass
}
