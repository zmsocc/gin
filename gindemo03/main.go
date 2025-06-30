package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	server := gin.Default()
	// 加载模板
	server.LoadHTMLGlob("templates/**/*")

	server.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好gin")
	})

	// 前台页面
	server.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"msg":   "我是 msg",
			"score": 90,
			"hobby": []string{"A", "B", "C"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻页面111",
					Content: "我是标题党111",
				},
				&Article{
					Title:   "新闻页面222",
					Content: "我是标题党222",
				},
			},
			"zms": map[string]string{
				"zzz": "lll",
				"mmm": "kkk",
				"sss": "jjj",
			},
			"test": []int{},
			"news": &Article{
					Title:   "新闻标题",
					Content: "新闻内容",
				},
		})
	})

	server.GET("/news", func(ctx *gin.Context) {
		news := &Article{
			Title:   "新闻页面",
			Content: "我是标题党",
		}
		ctx.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "这是一个新闻页面",
			"news":  news,
		})
	})

	// 后台页面
	server.GET("/admin/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	server.GET("/admin/news", func(ctx *gin.Context) {
		news := &Article{
			Title:   "新闻页面",
			Content: "我是标题党",
		}
		ctx.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "这是后台新闻页面",
			"news":  news,
		})
	})

	server.Run(":8080")
}
