package routers

import (

	"github.com/gin-gonic/gin"
	"github.com/zmsocc/gin/gindemo08/gindemo08_3/controllers/occ"
)

func DefaultRoutersInit(server *gin.Engine) {
	defaultRouters := server.Group("/index")
	{
		defaultRouters.GET("", occ.DefaultController{}.Index)
		defaultRouters.GET("/user", occ.DefaultController{}.News)
	}
}