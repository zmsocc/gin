package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo19/controllers/admin"
	"github.com/zmsocc/gin/gindemo19/middlewares"
)

func AdminRouters(ctx *gin.Engine) {
	adminRouters := ctx.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "后台首页")
		})

		adminRouters.GET("/article", admin.ArticleController{}.ArticleIndex)
		adminRouters.GET("/article/add", admin.ArticleController{}.ArticleAdd)
		adminRouters.GET("/article/edit", admin.ArticleController{}.ArticleEdit)
		adminRouters.GET("/article/delete", admin.ArticleController{}.ArticleDelete)

		adminRouters.GET("/user", admin.UserController{}.UserIndex)
		adminRouters.GET("/user/add", admin.UserController{}.UserAdd)
		adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)

		adminRouters.GET("/user/edit", admin.UserController{}.UserEdit)
		adminRouters.POST("/user/doedit", admin.UserController{}.DoEdit)

		adminRouters.GET("/user/edit02", admin.UserController{}.UserEditV2)
		adminRouters.POST("/user/doedit02", admin.UserController{}.DoEditV2)

		adminRouters.GET("/user/delete", admin.UserController{}.UserDelete)

		adminRouters.GET("/index", admin.IndexController{}.IndexContro)

		adminRouters.GET("/nav", admin.NavController{}.NavIndex)
		adminRouters.GET("/nav/add", admin.NavController{}.NavAdd)
		adminRouters.GET("/nav/edit", admin.NavController{}.NavEdit)
		adminRouters.GET("/nav/delete", admin.NavController{}.NavDelete)
	
		adminRouters.GET("/bank", admin.BankController{}.BankIndex)
	}
}