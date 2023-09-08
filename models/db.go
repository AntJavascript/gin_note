package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConDB() {
	dsn := "root:shi465608@tcp(127.0.0.1:3306)/note_app?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
	} else {
		fmt.Println("数据库连接成功", DB)
	}
}

func CloseDB() {
	DB.Close()
}
