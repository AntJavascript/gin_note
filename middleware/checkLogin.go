package middleware

import (
	"fmt"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("登录验证中间件")
		// 放行设置
		urlItem := []string{"/register", "/login"}
		if !tools.InStringArray(ctx.Request.RequestURI, urlItem) {
			// 从请求头中获取Token
			token := ctx.GetHeader("Authorization");
			_, err := tools.ParseToken(token)
			if err != nil {
				ctx.JSON(http.StatusOK, tools.JsonReturn("", "Token已过期", 401))
				ctx.Abort()
				return
			}
		}
		// 前置中间件
		ctx.Next()
	}
}
