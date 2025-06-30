package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	server := gin.Default()

	// 自定义模板函数 注意要把这个函数放在模板前
	server.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Print": Pri,
	})
	// 加载模板
	server.LoadHTMLGlob("templates/**/*")

	// 配置静态 web 目录， 第一个参数表示路由，第二个参数表示映射的目录
	server.Static("/static", "./static")

	server.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好gin")
	})

	// 前台页面
	server.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "aaa",
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
			"date": 1751294204,
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

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Pri(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "===" + str2
}
