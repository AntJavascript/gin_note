package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库配置
const (
	host     = "127.0.0.1"
	port     = "3306"
	username = "root"
	pwd      = "shi465608"
	database = "note_app"
)

var DB *gorm.DB

func init() {
	fmt.Println("初始化db", DB)
	conn := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	DB = db
	
	//自动创建数据表User
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Record{})
	DB.AutoMigrate(&model.Budget{})
}
