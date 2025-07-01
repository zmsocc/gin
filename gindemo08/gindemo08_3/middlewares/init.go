package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(ctx *gin.Context) {
	// 判断用户是否登录
	fmt.Println(time.Now().UnixMilli())
	fmt.Println(ctx.Request.URL)

	ctx.Set("username", "zhangsan")

	// 定义一个 goroutine 统计日志
	cCp := ctx.Copy()
	go func()  {
		time.Sleep(time.Second * 5)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}