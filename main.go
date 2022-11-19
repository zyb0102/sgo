package main

import (
	"database/sql"
	"fmt"
	"sgo/global"
	"sgo/initialize"
	"sgo/pkg/logs"
	"sgo/serve"
)

func main() {
	logs.Info("11111")
	logs.Warn("222222222222")
	logs.Error("11111")
}

func main2() {
	// 1.使用viper包加载配置
	global.Viper = initialize.InitViper()
	// 2.初始化日志库使用zap日志库
	global.Logger = initialize.InitZap()
	// 3.初始化数据库gorm库
	global.DB = initialize.InitGorm()
	// 程序结束关闭数据库
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
			fmt.Println("数据库连接已关闭")
		}(db)
	}
	// 4.初始化Redis服务
	if global.Config.App.OpenRedis {
		redis, err := initialize.InitRedis()
		if err != nil {
			global.Logger.Error(err.Error())
			return
		} else {
			global.Redis = redis
			defer global.Redis.Close()
		}
	}
	// 5.定时器
	// 6.运行服务
	serve.Run()
}
