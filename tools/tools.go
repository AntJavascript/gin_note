package tools

import (
	"fmt"
	"strings"
	"time"
)

// 日期格式转换
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

// 统一处理返回json数据
func JsonReturn(data interface{}, msg interface{}, code int) map[string]interface{} {
	res := make(map[string]interface{})

	res["data"] = data
	res["msg"] = msg
	res["code"] = code

	// 新的token
	access_token, refresh_token := GetNewToken(ctx)
	res["access_token"] = access_token
	res["refresh_token"] = refresh_token

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
func GetNewToken(ctx *gin.Context) (string, string){
	access_token := ctx.MustGet("access_token").(string)
	refresh_token := ctx.MustGet("refresh_token").(string)

	return access_token, refresh_token
}
