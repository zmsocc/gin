package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo10/models"
	"github.com/zmsocc/gin/gindemo10/routers"
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

	routers.AdminRouters(server)
	routers.DefaultRouters(server)

	server.Run(":8080")
}

