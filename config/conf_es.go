package config

import "fmt"

type ES struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	ArticleIndex  string `yaml:"article_index"`
	FullTextIndex string `yaml:"full_text_index"`
}

func (this ES) URL() string {
	return fmt.Sprintf("%s:%d", this.Host, this.Port)
}
