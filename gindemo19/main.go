package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo19/models"
	"github.com/zmsocc/gin/gindemo19/routers"
	"gopkg.in/ini.v1"
)

func main() {
	// 创建一个默认的路由引擎
	server := gin.Default()
	// 自定义模板函数，这个要放在加载模板前面
	server.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"GetDay": models.GetDay,
		"GetDate": models.GetDate,
	})
	// 加载模板，放在配置路由前面
	server.LoadHTMLGlob("templates/**/*")
	// 配置静态web目录，第一个参数表示路由，第二个参数表示映射的目录
	server.Static("/static", "./static")

	// 配置 session 中间件
	// stroe = cookie.NewStore("username")
	stroe, _ := redis.NewStore(10, "tcp", "localhost:6379", "张三","", []byte("secret"))
	server.Use(sessions.Sessions("mysession", stroe))


	routers.AdminRouters(server)
	routers.DefaultRouters(server)

	// 演示 "gopkg.in/ini.v1" 模块的使用
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 获取 ini 里面的配置
	fmt.Println(config.Section("").Key("app_name").String())
	fmt.Println(config.Section("mysql").Key("ip").String())
	fmt.Println(config.Section("redis").Key("ip").String())

	// config.Section("").Key("app_name").SetValue("你好 gin")
	// config.Section("").Key("admin_path").SetValue("/admin")
	// config.Section("mysql").Key("ip").SetValue("localhost")
	// config.SaveTo("./conf/app.ini")

	server.Run(":8080")
}

