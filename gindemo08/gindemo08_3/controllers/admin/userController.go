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
	username, _ := ctx.Get("username")
	fmt.Println(username)
	c.BaseController.Success(ctx)
	v, ok := username.(string)
	if ok {
		ctx.String(http.StatusOK, "用户列表---" + v)
	} else {
		ctx.String(http.StatusOK, "用户列表获取失败")
	}
}

func (c UserController) UserAdd(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户列表---Add")
}

func (c UserController) UserEdit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户列表---edit")
}