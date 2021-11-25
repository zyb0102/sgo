package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExampleApi struct {

}

func (a *ExampleApi) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,"11111111")
}
