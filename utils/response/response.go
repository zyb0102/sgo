package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 200
	ERROR   = 0
)

// Result 结果返回
func Result(code int, msg string, data interface{}, c *gin.Context) {
	// 响应结果
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// Success 成功返回结果
func Success(msg string, data interface{}, ctx *gin.Context)  {
	Result(SUCCESS,msg,data,ctx)
}

// Failed 失败返回结果
func Failed(msg string,ctx *gin.Context)  {
	Result(ERROR,msg, map[string]string{},ctx)
}