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
		total.GET("/year/:year", controllers.TotalStruct.GetYear)
	}

	// 预算相关
	budget := router.Group("budget")
	{
		budget.GET("/getBudget", controllers.BudgetStruct.GetBudget)
		budget.POST("/setBudget", controllers.BudgetStruct.SetBudget)
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
	
	/* 测试路由 */
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg": "测试成功",
		})
	})
	
	// 启动
	router.Run(":8080")
}
