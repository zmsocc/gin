package models

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Println("加载配置失败")
		os.Exit(1)
	}
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	database := config.Section("mysql").Key("database").String()
	
	// 用户名:密码@tcp(主机地址:端口)/要操作的数据库名字？
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",user, password, ip, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true, // 打印 sql
		// SkipDefaultTransaction: true, // 禁用事务 
	})
	if err != nil {
		fmt.Println(err)
	}
}
