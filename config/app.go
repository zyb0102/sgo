package config

type App struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
	Host string `mapstructure:"host" json:"host'" yaml:"host"`
	Port string `mapstructure:"port" json:"port'" yaml:"port"`
	OpenRedis bool `mapstructure:"open-redis" json:"open_redis'" yaml:"open-redis"`
}


