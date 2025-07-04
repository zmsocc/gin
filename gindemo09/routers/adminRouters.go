package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo09/controllers/admin"
	"github.com/zmsocc/gin/gindemo09/middlewares"
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

		adminRouters.GET("/user", admin.UserController{}.UserIndex)
		adminRouters.GET("/user/add", admin.UserController{}.UserAdd)
		adminRouters.GET("/user/edit", admin.UserController{}.UserEdit)

		adminRouters.GET("/index", admin.IndexController{}.IndexContro)
	}
}