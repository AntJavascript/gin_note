package router

import (
	"gin_note/controller"
	"gin_note/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
  // 初始化
	router := gin.Default()
  // 登录验证中间件
	router.Use(middleware.CheckLogin())
  // 设置静态资源路由
	router.Static("/static", "../static/img")

  /* 收支记录 */
	user := router.Group("record")
	{
		user.GET("/list", controller.Record.List)
		user.GET("/detail/:id", controller.Record.Detail)
		user.POST("/add", controller.Record.Add)
		user.POST("/update", controller.Record.Update)
		user.POST("/delete/:id", controller.Record.Delete)
	}
  
  // 启动
	router.Run(":9099")
}
