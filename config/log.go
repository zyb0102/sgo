package config

type Log struct {
	Directory string `mapstructure:"directory" json:"directory" yaml:"directory"`  // 日志存放目录
	Level string `mapstructure:"level" json:"level" yaml:"level"` // 写入日志等级
}
