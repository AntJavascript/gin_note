package main

import (
	_ "gin_note/router"
)

func main() {
	// 开始调试模式
	gin.SetMode("debug")
}
