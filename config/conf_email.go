package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             string `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` //发件人邮箱
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` //默认发件人名字
	UserSSL          bool   `json:"user_ssl" yaml:"user_ssl"`                     //是否使用SSL
	UserTLS          bool   `json:"user_tls" yaml:"user_tls"`
}
