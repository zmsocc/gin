package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(ctx *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(ctx.Request.URL)

	ctx.Set("username", "张三")

	cCp := ctx.Copy()
	go func ()  {
		time.Sleep(time.Second * 5)
		fmt.Println("Done! in path", cCp.Request.URL.Path)
	}()
}