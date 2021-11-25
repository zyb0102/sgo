package config

type Cors struct {
	// 开启域名(ip)限制
	OpenLimitOrigin bool `mapstructure:"open-limit-origin" json:"open_limit_origin" yaml:"open-limit-origin"`
	// 域名(ip)白名单
	WhiteOrigins []string `mapstructure:"white-origins" json:"white_origins" yaml:"white-origins"`
	// 域名(ip)黑名单
	BlackOrigins []string `mapstructure:"black-origins" json:"black_origins" yaml:"black-origins"`
	// 允许的header
	AllowHeaders []string `mapstructure:"allow-headers" json:"allow_headers" yaml:"allow-headers"`
}
