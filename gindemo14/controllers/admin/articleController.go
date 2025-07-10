package admin

import (
	"fmt"
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
	// article := []models.Article{}
	// models.DB.Where("id > ?", 2).Find(&article)
	article := models.Article{
		Id: 10,
		Title: "测试",
		CateId: 3,
		State: 1,
	}
	models.DB.Create(&article)
	fmt.Println(article)
	ctx.JSON(http.StatusOK, "添加文章成功")
}

func (c ArticleController) ArticleEdit(ctx *gin.Context) {
	article := models.Article{}
	models.DB.Model(&article).Where("id = ?", 10).Update("title", "啦啦啦")
	fmt.Println(article)
	ctx.String(http.StatusOK, "编辑文章成功")
}

func (c ArticleController) ArticleDelete(ctx *gin.Context) {
	article := models.Article{}
	models.DB.Where("id = ?", 10).Delete(&article)
	fmt.Println(article)
	ctx.String(http.StatusOK, "删除成功")
}