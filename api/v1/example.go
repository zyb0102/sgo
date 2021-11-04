package v1

import (
	"gin-api/utils"
	"gin-api/utils/response"
	"github.com/gin-gonic/gin"
)

type ExampleApi struct {
	
}

func (a *ExampleApi) Hello(ctx *gin.Context)  {
	response.Success("你好", utils.MD5("123456"),ctx)
}

