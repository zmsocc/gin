package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "值：%v" ,"Hello, World!")
	})
	server.GET("/news", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "我是新闻页面")
	})
	server.POST("/add", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "这是一个POST,主要用于增加数据")
	})
	server.PUT("/edit", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "这是一个PUT,主要用于编辑数据")
	})
	server.DELETE("/delete", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "这是一个DELETE,主要用于删除数据")
	})
	// 启动一个 web 服务器，监听在 8080 端口
	server.Run(":8080")
}