package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisOpt struct {
	Host     string // 主机ip
	Port     string // 端口号
	Username string // 用户名
	Password string // 密码
	DB       int    // 数据库
}

func NewRedis(o *RedisOpt) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", o.Host, o.Port),
		Username: o.Username,
		Password: o.Password,
		DB:       o.DB,
	})
	pong, err := cli.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("redis连接失败===>", err)
		return nil, err
	} else {
		fmt.Println(pong)
		return cli, nil
	}
}
