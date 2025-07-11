package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo17/models"
)

type NavController struct {}

func (c NavController) NavIndex(ctx *gin.Context) {
	navList := []models.Nav{}
	// models.DB.Find(&navList)
	// c.BaseController.Success(ctx)

	// 使用 in 查询
	// models.DB.Where("id in (?)", []int{1, 2, 4}).Find(&navList)

	// 使用 like 进行模糊查询
	// models.DB.Where("title like ?", "%会%").Find(&navList)

	// 使用 between and 进行范围查询
	// models.DB.Where("id between ? and ?", 3, 6).Find(&navList)

	// Or 或者查询
	// models.DB.Where("id = ? or id = ?", 2, 10).Find(&navList)

	// 使用 Select 指定返回的字段
	// models.DB.Select("id, title").Find(&navList)

	// Order排序，Limit，Offset
	// models.DB.Order("id desc").Limit(2).Offset(2).Find(&navList)

	// Count 统计总数
	// var num int64
	// models.DB.Find(&navList).Count(&num)

	// 使用原生 sql 删除表中的一条数据
	// models.DB.Exec("delete from nav where id = ?", 6)

	// 使用原生 sql 更改表中的一条数据
	// models.DB.Exec("update nav set title = '张三' where id = ?", 21)

	// // 使用原生 sql 查找表中的一条数据
	// models.DB.Raw("select * from nav where id = ?", 21).Scan(&navList)

	// 使用原生 sql 查找表中的一条数据
	// models.DB.Raw("select * from nav").Scan(&navList)
	
////// Exec 用于插入，删除，修改数据 //////
////// Raw 用于查询获取数据 //////

	// 使用原生 sql 统计表的行数
	var num int
	models.DB.Raw("select count(*) from nav").Scan(&num)
	fmt.Println(num)
	ctx.JSON(http.StatusOK,gin.H{
		"result": navList,
	})
}

func (c NavController) NavAdd(ctx *gin.Context) {
	// Nav := []models.Nav{}
	// models.DB.Where("id > ?", 2).Find(&Nav)
	nav := models.Nav{
		Id: 10,
		Title: "测试",
		Url: "www.zms.com",
		Status: 1,
		Sort: 1,
	}
	models.DB.Create(&nav)
	fmt.Println(nav)
	ctx.JSON(http.StatusOK, "添加Nav成功")
}

func (c NavController) NavEdit(ctx *gin.Context) {
	nav := models.Nav{}
	models.DB.Model(&nav).Where("id = ?", 10).Update("title", "啦啦啦")
	fmt.Println(nav)
	ctx.String(http.StatusOK, "编辑Nav成功")
}

func (c NavController) NavDelete(ctx *gin.Context) {
	nav := models.Nav{}
	models.DB.Where("id = ?", 10).Delete(&nav)
	fmt.Println(nav)
	ctx.String(http.StatusOK, "删除成功")
}