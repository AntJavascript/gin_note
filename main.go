package main

import (
	"example.com/m/v2/models"

	_ "example.com/m/v2/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 开始调试模式
	gin.SetMode("debug")

	models.ConDB()
}
