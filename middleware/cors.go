package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sgo/global"
	"sgo/utils"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		cors := global.Config.Cors
		// 跨域 域名处理
		origin := c.Request.Header.Get("origin")
		// 跨域解决
		if origin == "" {
			c.Header("Access-Control-Allow-Origin", "*")
		} else {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		// 跨域header处理,设置允许通过的header
		if len(cors.AllowHeaders) == 0 {
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		} else {
			allowHeaders := strings.Replace(strings.Trim(fmt.Sprint(cors.AllowHeaders), "[]"), " ", ",", -1)
			c.Header("Access-Control-Allow-Headers", allowHeaders)
		}
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 判断是否开启访问源限制限制
		if cors.OpenLimitOrigin {
			// 如果origin不在白名单里,或者在黑名单里禁止访问
			if !utils.IsValInList(origin, cors.WhiteOrigins) || utils.IsValInList(origin, cors.BlackOrigins) {
				c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{
					"code": "403",
					"msg":  "Forbidden",
				})
				return
			}
		}
		// 处理请求
		c.Next()
	}
}
