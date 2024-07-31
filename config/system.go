package config

type SystemConfig struct {
	Mode          string `mapstructure:"mode" json:"mode"`
	UrlPathPrefix string `mapstructure:"url-path-prefix" json:"urlPathPrefix"`
	Port          int    `mapstructure:"port" json:"port"`
}
