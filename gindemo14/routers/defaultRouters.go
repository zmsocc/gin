package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo14/controllers/occ"
	"github.com/zmsocc/gin/gindemo14/middlewares"
)

func DefaultRouters(ctx *gin.Engine) {
	DefaultController := ctx.Group("/index", middlewares.InitMiddleware)
	{
		DefaultController.GET("/", occ.DefaultController{}.Index)
		DefaultController.GET("/news", occ.DefaultController{}.News)
		DefaultController.GET("/deleteCookie", occ.DefaultController{}.DeleteCookie)
	}
}