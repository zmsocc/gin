package routers

import (

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo06/controllers/occ"
)

func DefaultRoutersInit(server *gin.Engine) {
	defaultRouters := server.Group("/index")
	{
		defaultRouters.GET("", occ.DefaultController{}.Index)
		defaultRouters.GET("/user", occ.DefaultController{}.News)
	}
}