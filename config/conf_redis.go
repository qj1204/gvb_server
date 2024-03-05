package config

import "fmt"

type Redis struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"pool_size" yaml:"pool_size"`
	TTL      int    `json:"ttl" yaml:"ttl"`
}

func (this *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", this.IP, this.Port)
}
