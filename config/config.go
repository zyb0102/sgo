package config

type Config struct {
	Log  // 日志相关配置
	Gorm // gorm 相关配置
	Mysql  // mysql相关配置
}
