package main

import (
	_ "gin_note/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 开始调试模式
	gin.SetMode("debug")
}
