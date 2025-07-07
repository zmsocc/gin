package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (c ArticleController) ArticleIndex(ctx *gin.Context) {
	c.BaseController.Success(ctx)
	ctx.String(http.StatusOK, "文章首页")
}

func (c ArticleController) ArticleAdd(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加文章---ADD")
}

func (c ArticleController) ArticleEdit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "编辑文章---Edit")
}