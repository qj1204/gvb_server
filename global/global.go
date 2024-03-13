package global

import (
	"github.com/olivere/elastic/v7"
	"github.com/oschwald/geoip2-golang"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/config"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface // MysqlLog显示所有的sql
	Redis    *redis.Client
	ESClient *elastic.Client
	AddrDB   *geoip2.Reader
)
