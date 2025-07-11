package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (c IndexController) IndexContro(ctx *gin.Context) {
	c.BaseController.Success(ctx)
	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "admin/index首页",
	})
}