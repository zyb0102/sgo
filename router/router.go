package router

import (
	v1 "gin-api/router/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	v1 v1.RouterGroup
}

func Register(r *gin.Engine) {
	var routerGroup Router
	router := r.Group("")
	{
		// 测试路由
		r.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,"hello,world")
		})
		// 注册v1路由
		routerGroup.v1.InitRouter(router)
	}
}
