package occ

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (c DefaultController) Index(ctx *gin.Context) {
	// fmt.Println(models.UnixToTime(1629788418))
	// fmt.Println(models.GetDay())
	// fmt.Println(models.GetUnix())

	// 设置 Cookie
	// 3600 秒
	ctx.SetCookie("username", "李四", 3600, "/index", "localhost", false, true)
	// ctx.SetCookie("hobby", "吃饭 睡觉", 10, "/index", "localhost", false, true)

	// 共享二级域名
	// ctx.SetCookie("username", "李四", 3600, "/index", ".occ.com", false, true)

	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"title": "我是default首页标题",
		"content": "我是default首页内容",
		"t": 1629788418,
		// "d": models.GetDay,
	})
}

func (c DefaultController) News(ctx *gin.Context) {
	// 获取 Cookie 
	username, _ := ctx.Cookie("username")
	// hobby, _ := ctx.Cookie("hobby")
	ctx.String(http.StatusOK, "username=%v",username)

	// ctx.HTML(http.StatusOK, "default/news.html", gin.H{
	// 	"title": "我是新闻首页标题",
	// 	"content": "我是新闻内容",
	// })
}

func (c DefaultController) DeleteCookie(ctx *gin.Context) {
	// 获取Cookie
	// 过期时间<0，就是删除Cookie
	ctx.SetCookie("username", "李四", -1, "/index", "localhost", false, true)
	ctx.String(http.StatusOK, "成功删除Cookie")
}