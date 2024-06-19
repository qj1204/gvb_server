package config

type Config struct {
	Mysql         Mysql         `yaml:"mysql"`
	LoggerSetting LoggerSetting `yaml:"logger_setting"`
	System        System        `yaml:"system"`
	Upload        Upload        `yaml:"upload"`
	SiteInfo      SiteInfo      `yaml:"site_info"`
	Email         Email         `yaml:"email"`
	QQ            QQ            `yaml:"qq"`
	QiNiu         QiNiu         `yaml:"qi_niu"`
	Jwt           Jwt           `yaml:"jwt"`
	Redis         Redis         `yaml:"redis"`
	ES            ES            `yaml:"es"`
	ChatGroup     ChatGroup     `yaml:"chat_group"`
	Gaode         Gaode         `yaml:"gaode"`
	BigModel      BigModel      `yaml:"big-model"`
}
