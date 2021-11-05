package config

type Cors struct {
	AccessControlAllowOrigin []string `mapstructure:"access-control-allow-origin" json:"access_control_allow_origin" yaml:"access-control-allow-origin"`
	AccessControlAllowHeaders []string `mapstructure:"access-control-allow-headers" json:"access_control_allow_headers" yaml:"access-control-allow-headers"`
}
