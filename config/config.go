package config

type Config struct {
	App // app相关配置
	Log  // 日志相关配置
	Gorm // gorm 相关配置
	Mysql  // mysql相关配置
	Cors // 跨域相关配置
}
