package config

type Mysql struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	UserName string `mapstructure:"user-name" json:"user_name" yaml:"user-name"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Charset string `mapstructure:"charset" json:"charset" yaml:"charset"`
	ParseTime string `mapstructure:"parse-time" json:"parse_time" yaml:"parse-time"`
	Loc string `mapstructure:"loc" json:"loc" yaml:"loc"`
}