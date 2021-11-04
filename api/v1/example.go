package v1

import (
	"gin-api/utils/response"
	"github.com/gin-gonic/gin"
)

type ExampleApi struct {
	
}

func (a *ExampleApi) hello(ctx *gin.Context)  {
	response.Success("你好", map[string]string{"ps":"初次见面"},ctx)
}

