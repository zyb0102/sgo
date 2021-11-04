package v1

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	ExampleRouterGroup  // example路由组
}

func (g *RouterGroup) InitRouter(Router *gin.RouterGroup) {
	// 路由设置为v1
	v1Router := Router.Group("v1")
	{
		// 初始化Example路由
		g.ExampleRouterGroup.InitRouter(v1Router)
	}
}