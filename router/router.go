package router

import (
	"example.com/m/v2/controllers"
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
	record := router.Group("record")
	{
		record.GET("/list", controllers.Record.List)
		record.GET("/detail/:id", controllers.Record.Detail)
		record.POST("/add", controllers.Record.Add)
		record.POST("/update", controllers.Record.Update)
		record.POST("/delete/:id", controllers.Record.Delete)
	}
	
	// 统计相关
	total := router.Group("total")
	{
		total.GET("/day", controllers.TotalStruct.GetDay)
		total.GET("/month", controllers.TotalStruct.GetMonth)
		// record.GET("/year/:year", controllers.TotalStruct.GetYear)
		// record.GET("/getBudget", controllers.TotalStruct.GetBudget)
		// record.post("/setBudget", controllers.TotalStruct.SetBudget)
	}
	
	/* 登录相关 */
	user := router.Group("user")
	{
		user.POST("/login", controllers.User.Login)
		user.POST("/register", controllers.User.Register)
	}
	
	/* jwt相关 */
	jwt := router.Group("jwt")
	{
		jwt.POST("/generateToken", controllers.JwtStruct.Generate)
		jwt.POST("/parseToken", controllers.JwtStruct.Parse)
	}
	
	// 启动
	router.Run(":9099")
}
