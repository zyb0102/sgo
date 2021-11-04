package v1

import (
	v1 "gin-api/api/v1"
	"github.com/gin-gonic/gin"
)

type ExampleRouterGroup struct {

}

func (g *ExampleRouterGroup) InitRouter(Router *gin.RouterGroup)  {
	var exampleApi v1.ExampleApi
	routerGroup := Router.Group("example")
	{
		routerGroup.GET("hello",exampleApi.Hello)
	}
}
