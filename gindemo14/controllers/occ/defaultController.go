package occ

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (c DefaultController) Index(ctx *gin.Context) {
	// 设置 sessions
	session := sessions.Default(ctx)
	session.Set("username", "zhangSan")
	// 配置 session 的过期时间
	// session.Options(sessions.Options{
	// 	MaxAge: 10,	// MaxAge 单位是秒
	// })
	// 必须调用 session.Save
	session.Save()
	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"title": "我是Index首页",
		"content": "我是Index内容",
	})
}

func (c DefaultController) News(ctx *gin.Context) {
	// 获取 session
	session := sessions.Default(ctx)
	username := session.Get("username")
	name, _ := username.(string)
	ctx.String(http.StatusOK, "session="+name)
}

func (c DefaultController) DeleteCookie(ctx *gin.Context) {
	// 获取Cookie
	// 过期时间<0，就是删除Cookie
	ctx.SetCookie("username", "李四", -1, "/index", "localhost", false, true)
	ctx.String(http.StatusOK, "成功删除Cookie")
}