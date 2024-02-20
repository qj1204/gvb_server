package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` // 高级配置，例如：charset
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级，debug就是输出全部sql（dev，release）
}

func (this *Mysql) Dsn() string {
	return this.User + ":" + this.Password + "@tcp(" + this.Host + ":" + fmt.Sprintf("%d", this.Port) + ")/" + this.DB + "?" + this.Config
}
