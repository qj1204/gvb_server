package config

import (
	"fmt"
)

type ES struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	ArticleIndex  string `yaml:"article_index"`
	FullTextIndex string `yaml:"full_text_index"`
}

func (es ES) URL() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
