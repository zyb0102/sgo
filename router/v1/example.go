package v1

import (
	"gin-api/application/v1/controller"
	"github.com/gin-gonic/gin"
)

type ExampleRouterGroup struct {

}

func (g *ExampleRouterGroup) InitRouter(Router *gin.RouterGroup)  {
	var api controller.ExampleApi
	router := Router.Group("example")
	{
		router.GET("hello",api.Hello)
	}
}
