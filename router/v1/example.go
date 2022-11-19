package v1

import (
	"github.com/gin-gonic/gin"
	"sgo/app/v1/controller"
)

type ExampleRouterGroup struct {
}

func (g *ExampleRouterGroup) InitRouter(Router *gin.RouterGroup) {
	var api controller.ExampleApi
	router := Router.Group("example")
	{
		router.GET("hello", api.Hello)
	}
}
