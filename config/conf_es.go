package config

import "fmt"

type ES struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (this ES) URL() string {
	return fmt.Sprintf("%s:%d", this.Host, this.Port)
}
