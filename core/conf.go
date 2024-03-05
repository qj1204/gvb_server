package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "settings.yaml"

// InitConf 读取settings.yaml配置
func InitConf() {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf errer: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		// 此时配置文件还没读取到，所以不能用自己的logrus
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFiles load Init success.")
	global.Config = c
}

// SetYaml 设置yaml配置文件
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
