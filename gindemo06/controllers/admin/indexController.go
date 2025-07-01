package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Index===")
}