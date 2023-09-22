package constants

import (
	"net/http"
)

const (
	// token日期
  ACCESSTOKEN  = 2 // 2小时
  REFRESHTOKEN = 720 // 30天

	// 返回msg提示语
	SUCCESS = "success"
	FAIL    = "fail"

	// 返回状态码
	STATUSOK = http.StatusOK
)

// 标准json返回函数参数
type CommonJson struct {
	Ctx    *gin.Context
	Data   interface{}
	Msg    string
	Code   int
	NewToken bool
}
