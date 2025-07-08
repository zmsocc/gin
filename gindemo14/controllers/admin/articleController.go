package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo14/models"
)

type ArticleController struct {
	BaseController
}

func (c ArticleController) ArticleIndex(ctx *gin.Context) {
	articleList := []models.Article{}
	models.DB.Find(&articleList)
	// c.BaseController.Success(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"result": articleList,
	})
}

func (c ArticleController) ArticleAdd(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加文章---ADD")
}

func (c ArticleController) ArticleEdit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "编辑文章---Edit")
}