package occ

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (d DefaultController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是一个=== msg ===",
	})
}

func (d DefaultController) News(ctx *gin.Context) {
	ctx.String(http.StatusOK, "News")
}
