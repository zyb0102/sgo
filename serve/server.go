package serve

import (
	"fmt"
	gzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"sgo/global"
	"sgo/middleware"
	"sgo/router"
	"time"
)

func Run() {
	r := gin.New()
	// 使用中间件将所有错误信息和请求信息记录到日志
	r.Use(gzap.Ginzap(global.Logger, time.RFC3339, true))
	r.Use(gzap.RecoveryWithZap(global.Logger, true))
	// 使用跨域中间件
	r.Use(middleware.Cors())

	// 路由注册start
	router.Register(r)
	// 路由注册结束
	appConfigParams := global.Config.App
	// 运行服务
	address := fmt.Sprintf("%s:%s", appConfigParams.Host, appConfigParams.Port)
	err := r.Run(address)
	if err != nil {
		fmt.Println("服务启动失败")
		return
	}
}
