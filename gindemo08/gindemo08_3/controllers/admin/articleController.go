package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (c ArticleController) ArticleIndex(ctx *gin.Context) {
	c.BaseController.Fail(ctx)
}

func (c ArticleController) ArticleAdd(ctx *gin.Context) {
	ctx.String(http.StatusOK, "---add--文章")
}

func (c ArticleController) ArticleEdit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "---edit--文章")
}