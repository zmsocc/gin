package main

import (
	"html/template"
	"time"

	"github.com/zmsocc/gin/gindemo06/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	server :=  gin.Default()

	// 自定义模板函数, 放在加载模板前面
	server.SetFuncMap(template.FuncMap{
		"UnixTotime": UnixTotime,
	})
	// 加载模板
	server.LoadHTMLGlob("templates/**/*")
	// 配置静态web路由
	server.Static("/static", "./static")


	routers.DefaultRoutersInit(server)
	routers.AdminRoutersInit(server)

	server.Run(":8080")
}

func UnixTotime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}