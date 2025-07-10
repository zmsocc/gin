package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo15/models"
)

type ArticleController struct {
	BaseController
}

func (c ArticleController) ArticleIndex(ctx *gin.Context) {
	// 一对一
	// // articleList := []models.Article{}
	// articleList := []models.ArticleV2{}
	// models.DB.Preload("ArticleC").Find(&articleList)
	// // c.BaseController.Success(ctx)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"result": articleList,
	// })

	// 一对多
	// articleCateList := []models.ArticleCate{}
	// models.DB.Preload("Article").Find(&articleCateList)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"result": articleCateList,
	// })

	// 多对多
	// studentList := []models.Student{}
	// models.DB.Preload("Les son").Find(&studentList)

	// 查询课程被哪些学生选修，排除张三
	lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "id != 1").Find(&lessonList)

	// 查询课程被哪些学生选修，排除张三,李四
	// models.DB.Preload("Student", "id not in (?)", []int{1, 2}).Find(&lessonList)
	
	// 查看课程被那些学生选修，要求：学生 id 倒叙输出，自定义预加载 SQL
	// models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
	// 	return models.DB.Order("student.id desc")
	// }).Find(&lessonList)
	ctx.JSON(http.StatusOK, gin.H{
		"result": lessonList,
	})
}

func (c ArticleController) ArticleAdd(ctx *gin.Context) {
	// article := []models.Article{}
	// models.DB.Where("id > ?", 2).Find(&article)
	article := models.Article{
		Id: 10,
		Title: "测试",
		ArticleCateId: 3,
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