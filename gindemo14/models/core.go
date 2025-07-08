package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// 用户名:密码@tcp(主机地址:端口)/要操作的数据库名字？
	dsn := "root:root@tcp(localhost:3306)/occ?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}