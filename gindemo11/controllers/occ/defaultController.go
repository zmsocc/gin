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
	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"title": "我是default首页标题",
		"content": "我是default首页内容",
		"t": 1629788418,
		// "d": models.GetDay,
	})
}

func (c DefaultController) News(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/news.html", gin.H{
		"title": "我是新闻首页标题",
		"content": "我是新闻内容",
	})
}