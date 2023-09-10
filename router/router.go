package router

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/models"

	// "example.com/m/v2/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化
	router := gin.Default()
	// 登录验证中间件
	// router.Use(middleware.CheckLogin())
	// 设置静态资源路由
	router.Static("/static", "../static/img")

	/* 收支记录 */
	user := router.Group("record")
	{
		user.GET("/list", controllers.Record.List)
		user.GET("/detail/:id", controllers.Record.Detail)
		// user.POST("/add", controllers.Record.Add)
		// user.POST("/update", controllers.Record.Update)
		// user.POST("/delete/:id", controllers.Record.Delete)
	}
	models.ConDB()
	// 启动
	router.Run(":9099")
}
