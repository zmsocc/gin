package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// 自定义模板，这个要放在加载模板前面
	server.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	// 加载模板
	server.LoadHTMLGlob("templates/**/*")
	

	server.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"date": 1751345915,
		})
	})

	server.GET("/news", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻首页",
		})
	})

	// Get 请求传值
	server.GET("/", func(ctx *gin.Context) {
		username := ctx.Query("username")
		age := ctx.Query("age")
		page := ctx.DefaultQuery("page", "1")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"age": age,
			"page": page,
		})
	})

	server.GET("/article", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "1")

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id": id,
		})
	})

	// POST 演示
	server.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/user.html", gin.H{})
	})

	server.POST("/doAddUser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		age := ctx.DefaultPostForm("age", "20")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age": age,
		})
	})

	// 获取 GET POST 传递的数据绑定到结构体
	server.GET("/getuser", func(ctx *gin.Context) {
		user := &UserInfo{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err,
			})
			return
		}
		fmt.Printf("%#v", user)
		ctx.JSON(http.StatusOK, user)
	})

	server.POST("/doAddUser2", func(ctx *gin.Context) {
		user := &UserInfo{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	// 获取 xml 数据
	server.POST("/xml", func(ctx *gin.Context) {
		article := &Article{}
		// 获取 ctx.Request.Body 读取请求数据
		xmlSliceData, _ := ctx.GetRawData()
		fmt.Println(xmlSliceData)
		err := xml.Unmarshal(xmlSliceData, &article)
		if err != nil {
			ctx.JSON(http.StatusOK, "xml解析错误")
		}
		ctx.JSON(http.StatusOK, article)
	})

	// 动态路由传值
	server.GET("/list/:cid", func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		ctx.String(http.StatusOK, "%v", cid)
	})

	server.Run(":8080")
}

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Age string 	`form:"age"`
}

type Article struct {
	Title string `xml:"title"`
	Content string `json:"content" xml:"content"`
}
