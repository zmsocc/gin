package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (c UserController) UserIndex(ctx *gin.Context) {
	c.BaseController.Success(ctx)
	name, ok := ctx.Get("username")
	fmt.Println(name)
	if !ok {
		ctx.String(http.StatusOK, "获取用户名失败")
		return
	}
	username := name.(string)
	ctx.String(http.StatusOK, "用户列表获取成功---"+username)
}

func (c UserController) UserAdd(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户===ADD")
}

func (c UserController) UserEdit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "编辑用户===Edit")
}