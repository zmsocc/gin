package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (c BaseController) Success(ctx *gin.Context) {
	ctx.String(http.StatusOK, "成功")
}

func (c BaseController) Fail(ctx *gin.Context) {
	ctx.String(http.StatusOK, "失败")
}