package initialize

import (
	"fmt"
	"gin-api/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取错误,config file: %s \n", err))
	}
	err = v.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("配置文件参数解析失败: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		global.Logger.Info("配置文件发生更新,重新加载配置文件")
		// 重新解析配置文件
		err = v.Unmarshal(&global.Config)
		if err != nil {
			panic(fmt.Errorf("配置文件参数重新解析失败: %s \n", err))
		}
	})


	return v
}
