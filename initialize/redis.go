package initialize

import (
	"context"
	"errors"
	"fmt"
	"gin-api/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)
// InitRedis 初始化redis
func InitRedis() (*redis.Client, error) {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s",redisCfg.Host,redisCfg.Port),
		Password: redisCfg.Password,
		DB:       redisCfg.Database,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("redis connect ping failed, err:", zap.Any("err", err))
		return client,errors.New("redis连接失败")
	} else {
		global.Logger.Info("redis connect ping response:", zap.String("pong", pong))
		return client, nil
	}
}
