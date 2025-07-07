package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo10/controllers/occ"
	"github.com/zmsocc/gin/gindemo10/middlewares"
)

func DefaultRouters(ctx *gin.Engine) {
	DefaultController := ctx.Group("/index", middlewares.InitMiddleware)
	{
		DefaultController.GET("/", occ.DefaultController{}.Index)
		DefaultController.GET("/news", occ.DefaultController{}.News)
	}
}