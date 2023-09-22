package tools

import (
	"fmt"
	"strings"
	"example.com/m/v2/constants"
	"github.com/gin-gonic/gin"
)

// 统一处理返回json数据
func JsonReturn(data *constants.CommonJson) map[string]interface{} {
	res := make(map[string]interface{})

	res["data"] = data.data
	res["msg"] = data.msg
	res["code"] = data.code

	// 是否生成新的token
	if(data.newToken) {
		access_token, refresh_token := GetNewToken(data.ctx)
		res["access_token"] = access_token
		res["refresh_token"] = refresh_token
	}
	
	return res
}

func Split(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

func Replace(origin, search, replace string, count ...int) string {
	n := -1
	if len(count) > 0 {
		n = count[0]
	}
	return strings.Replace(origin, search, replace, n)
}

func Equal(a, b string) bool {
	return strings.EqualFold(a, b)
}

func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

func SubStr(str string, start int, length ...int) (substr string) {
	lth := len(str)

	// Simple border checks.
	if start < 0 {
		start = 0
	}
	if start >= lth {
		start = lth
	}
	end := lth
	if len(length) > 0 {
		end = start + length[0]
		if end < start {
			end = lth
		}
	}
	if end > lth {
		end = lth
	}
	return str[start:end]
}

func Join(array []string, sep string) string {
	return strings.Join(array, sep)
}

func UcWords(str string) string {
	return strings.Title(str)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// 获取新的token
func GetNewToken(ctx *gin.Context) (string, string) {
	access_token := ctx.MustGet("access_token").(string)
	refresh_token := ctx.MustGet("refresh_token").(string)

	return access_token, refresh_token
}

// 数字转字符串
func ToString(value int) string {
	return fmt.Sprintf("%s", value)
}

// 获取请求的头部token相关信息
func GetHearToken(ctx *gin.Context) (string, string) {
	access_token := ctx.Request.Header.Get("access_token")
	refresh_token := ctx.Request.Header.Get("refresh_token")

	return access_token, refresh_token
}
