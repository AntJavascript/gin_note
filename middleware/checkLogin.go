package middleware

import (
	"fmt"
	"net/http"

	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
)

// 设置token
func setToken(ctx *gin.Context, access_token, refresh_token string) {
	ctx.Set("access_token", access_token)
	ctx.Set("refresh_token", refresh_token)
}

func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("登录验证中间件")
		// 放行设置
		urlItem := []string{"/user/register", "/user/login", "/jwt/generateToken", "/jwt/parseToken"}
		if !tools.InStringArray(ctx.Request.RequestURI, urlItem) {
			// 从请求头中获取Token
			access_token := ctx.GetHeader("access_token")
			refresh_token := ctx.GetHeader("refresh_token")
			if len(access_token) == 0 || len(refresh_token) == 0 {
				setToken(ctx, "", "")
				ctx.JSON(http.StatusOK, tools.JsonReturn(ctx, "", "鉴权失败", 401))
				ctx.Abort()
				return
			}

			// 验证access_token
			_, err := tools.ParseToken(access_token)
			// 验证refresh_token
			_, refresh_err := tools.ParseToken(refresh_token)

			// access_token已过期、refresh_token未过期, 重新生成token
			if err != nil && refresh_err == nil {
				new_access_token, new_refresh_token, _ := tools.RefreshToken(access_token, refresh_token)
				setToken(ctx, new_access_token, new_refresh_token)
			} else if refresh_err != nil {
				setToken(ctx, "", "")
				ctx.JSON(http.StatusOK, tools.JsonReturn(ctx, "", "Token已过期, 请重新登录", 401))
				ctx.Abort()
				return
			}

		} else {
			setToken(ctx, "", "") // 放行设置，也需要设置token为空
		}
		setToken(ctx, "", "")
		// 前置中间件
		ctx.Next()
	}
}
