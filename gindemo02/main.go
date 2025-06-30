package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserList struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Content string `json:"content666"`
}

func main() {
	server := gin.Default()
	// 配置模板的文件
	server.LoadHTMLGlob("templates/*")

	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	server.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "Hello, JSON!",
		})
	})
	server.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "Hello, JSON2!",
		})
	})
	// :8080/jsonp?callback=abc
	// 会返回abc({"success":true,"msg":"Hello, JSONP!"})
	server.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"success": true,
			"msg":     "Hello, JSONP!",
		})
	})
	server.GET("/json3", func(c *gin.Context) {
		c.JSON(http.StatusOK, &UserList{
			Title:   "User List",
			Desc:    "List of users",
			Content: "Here are the users",
		})
	})

	server.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "Hello, XML!",
		})
	})

	
	server.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品页面",
			"price":  100,
		})
	})
	server.Run(":8080")
}