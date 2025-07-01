package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"


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

	server.GET("/index", initMiddleware, initMiddlewareV2, func(ctx *gin.Context) {
		fmt.Println("这是我的首页")
		time.Sleep(time.Second)
		ctx.String(http.StatusOK, "gin 首页")
	})

	server.GET("/news", initMiddleware, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "新闻页面")
	})

	server.Run(":8080")
}

func UnixTotime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func initMiddleware(ctx *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件")

	// 调用该请求的剩余处理程序
	ctx.Next()

	// 终止后面的请求
	// ctx.Abort() 

	fmt.Println("2-我是一个中间件")
	end := time.Now().UnixNano()

	fmt.Println(end-start)
}

func initMiddlewareV2(ctx *gin.Context) {
	fmt.Println("1-我是一个中间件-V2")

	// 调用该请求的剩余处理程序
	ctx.Next()

	// 终止后面的请求
	// ctx.Abort() 

	fmt.Println("2-我是一个中间件-V2")
}