package minddleware

import (
	"fmt"
	"gin-api/global"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		corsConfigParams := global.Config.Cors
		// 跨域处理 origin
		fmt.Println(corsConfigParams.AccessControlAllowOrigin)
		fmt.Println(len(corsConfigParams.AccessControlAllowOrigin))
		if len(corsConfigParams.AccessControlAllowOrigin) == 0 {
			c.Header("Access-Control-Allow-Origin", "*")
		} else {
			origin := c.Request.Header.Get("origin")
			fmt.Println(origin)
			if origin != "" && utils.IsValInList(origin,corsConfigParams.AccessControlAllowOrigin) {
				c.Header("Access-Control-Allow-Origin", origin)
			}
		}
		// 跨域header处理
		if len(corsConfigParams.AccessControlAllowHeaders) == 0 {
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		} else {
			accessControlAllowHeaders := strings.Replace(strings.Trim(fmt.Sprint(corsConfigParams.AccessControlAllowHeaders), "[]"), " ", ",", -1)
			c.Header("Access-Control-Allow-Headers", accessControlAllowHeaders)
		}


		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
