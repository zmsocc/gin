package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo08/gindemo08_3/controllers/admin"
	"github.com/zmsocc/gin/gindemo08/gindemo08_3/middlewares"
)


func AdminRoutersInit(server *gin.Engine) {
	adminRouters := server.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "后台首页")
		})
		adminRouters.GET("/user", admin.UserController{}.UserIndex)
		adminRouters.GET("/user/add", admin.UserController{}.UserAdd)
		adminRouters.GET("/user/edit", admin.UserController{}.UserEdit)

		adminRouters.GET("index")

		adminRouters.GET("/article", admin.ArticleController{}.ArticleIndex)
		adminRouters.GET("/article/add", admin.ArticleController{}.ArticleAdd)
		adminRouters.GET("/article/edit", admin.ArticleController{}.ArticleEdit)
	}
}