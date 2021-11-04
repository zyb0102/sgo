package serve

import (
	"fmt"
	"gin-api/global"
	"gin-api/router"
	gzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func Run() {
	r := gin.New()
	// 使用中间件将所有错误信息和请求信息记录到日志
	r.Use(gzap.Ginzap(global.Logger, time.RFC3339, true))
	r.Use(gzap.RecoveryWithZap(global.Logger, true))

	// 路由注册start
	router.Register(r)
	// 路由注册结束
	// 运行服务
	err := r.Run("192.168.10.64:9001")
	if err != nil {
		fmt.Println("服务启动失败")
		return
	}
}
