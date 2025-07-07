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
	ctx.String(http.StatusOK, "这是首页")
}