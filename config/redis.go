package config

type Redis struct {
	// 数据库名
	Database   int    `mapstructure:"database" json:"database" yaml:"database"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
