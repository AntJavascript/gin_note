package controllers

import (
	"net/http"
	"example.com/m/v2/dto"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
)

var JwtStruct = new(jwtCtrl)

type jwtCtrl struct{
	Token string
}

// 生成token
func (c *jwtCtrl) Generate(ctx *gin.Context) {
	res := make(map[string]interface{})

	// 解析前端body参数
	params := dto.User{}
	ctx.BindJSON(&params)

	token, err := tools.GenerateToken(params.Phone, params.Password, 2)

	if err != nil {
		res = tools.JsonReturn(err, "生成token失败", 400)
	} else {
		res = tools.JsonReturn(err, "生成token成功",200)
	}
	res["token"] = token
	ctx.JSON(http.StatusOK, res)
	return
}

// 解析token
func (c *jwtCtrl) Parse(ctx *gin.Context) {
	res := make(map[string]interface{})

	// 解析前端body参数
	params :=jwtCtrl{}
	ctx.BindJSON(&params)

	claims, err := tools.ParseToken(params.Token)

	if err != nil {
		res = tools.JsonReturn(err, "解析token失败", 400)
	} else {
		res = tools.JsonReturn(err, "解析token成功",200)
	}
	res["claims"] = claims
	ctx.JSON(http.StatusOK, res)
	return
}
