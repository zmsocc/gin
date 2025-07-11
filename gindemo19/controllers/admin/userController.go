package admin

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo19/models"
)

type UserController struct {
	BaseController
}

func (c UserController) UserIndex(ctx *gin.Context) {
	// userList := []models.User{}
	// models.DB.Find(&userList)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"result": userList,
	// })

	userList := []models.User{}
	models.DB.Where("age>25").Find(&userList)
	ctx.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

func (c UserController) UserAdd(ctx *gin.Context) {
	user := models.User{
		Id: 8,
		Username: "zms666",
		Email: "zxms@qq.com",
		Age: 38,
		AddTime: time.Now().Unix(),
	}
	models.DB.Create(&user)
	fmt.Println(user)
	ctx.String(http.StatusOK, "增加数据成功")
}

func (c UserController) UserEdit(ctx *gin.Context) {
	// 查询 id 等于 7 的数据
	// user := models.User{
	// 	Id: 7,
	// }
	// models.DB.Find(&user)
	// // 更新数据
	// user.Username = "张三"
	// user.Email = "001@qq.com"
	// user.AddTime = time.Now().Unix()
	user := models.User{}
	models.DB.Model(&user).Where("id = ?", 7).Update("username","王五")

	// models.DB.Save(&user)
	fmt.Println(user)
	ctx.String(http.StatusOK, "修改成功")
}

func (c UserController) UserDelete(ctx *gin.Context) {
	// user := models.User{Id: 9}
	// models.DB.Delete(&user)
	models.DB.Where("id = ?", 8).Delete(&models.User{})
	ctx.String(http.StatusOK, "删除成功")
}

func (c UserController) DoUpload(ctx *gin.Context) {
	name := ctx.PostForm("username")
	file, err := ctx.FormFile("face")
	if err != nil {
		ctx.String(http.StatusOK, "获取文件失败")
		return
	}
	fileTailMap := map[string]bool{
		".jpg": true,
		".png": true,
		".gif": true,
		".jpeg": true,
	}
	extName := path.Ext(file.Filename)
	if _, ok := fileTailMap[extName]; !ok {
		ctx.String(http.StatusOK, "图片不合法")
		return
	}

	// 创建图片保存目录 
	day := models.GetDay()
	dir := "./static/upload/" + day
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		ctx.String(http.StatusOK, "MkdirAll失败")
		return
	}
	// 生成文件保存目录
	fileName := strconv.FormatInt(models.GetUnix(), 10) + extName
	// 执行上传
	dst := path.Join(dir, fileName)
	ctx.SaveUploadedFile(file, dst)
	ctx.JSON(http.StatusOK, gin.H{
		"success":true,
		"username": name,
	})
	// // 获取用户名
	// name := ctx.PostForm("username")
	// // 获取文件
	// file, err := ctx.FormFile("face")
	// if err != nil {
	// 	ctx.String(http.StatusOK, "获取文件失败")
	// 	return
	// }
	// // 路径拼接
	// dst := path.Join("./static/upload", file.Filename)
	// // 文件上传至指定路径
	// ctx.SaveUploadedFile(file, dst)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// 	"username": name,
	// 	"dst": dst,
	// })
}

func (c UserController) DoEdit(ctx *gin.Context) {
	// 获取用户名
	name := ctx.PostForm("username")
	// 获取文件
	file1, err1 := ctx.FormFile("face1")
	if err1 != nil {
		ctx.String(http.StatusOK, "获取文件失败")
		return
	}
	// 路径拼接
	dst1 := path.Join("./static/upload", file1.Filename)
	// 文件上传至指定路径
	ctx.SaveUploadedFile(file1, dst1)

	file2, err2 := ctx.FormFile("face2")
	if err2 != nil {
		ctx.String(http.StatusOK, "获取文件失败")
		return
	}
	dst2 := path.Join("./static/upload", file2.Filename)
	ctx.SaveUploadedFile(file2, dst2)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"username": name,
		"dst1": dst1,
		"dst2": dst2,
	})
}

func (c UserController) UserEditV2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/useredit02.html", gin.H{
		"title": "添加用户---ADD",
	})
}

func (c UserController) DoEditV2(ctx *gin.Context) {
	// 获取用户名
	name := ctx.PostForm("username")
	// 获取文件
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusOK, "获取文件失败")
		return
	}
	files := form.File["face[]"]
	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		ctx.SaveUploadedFile(file, dst)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"username": name,
	})
}