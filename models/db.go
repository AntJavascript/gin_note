package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库配置
const (
	host = 127.0.0.1
	port = 3306
	username = root
	pwd = shi465608
	database = note_app
)

var DB *gorm.DB

func ConDB() {
	conn := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
	} else {
		fmt.Println("数据库连接成功", DB)
	}
}

func CloseDB() {
	DB.Close()
}
