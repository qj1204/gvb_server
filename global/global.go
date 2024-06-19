package global

import (
	"github.com/olivere/elastic/v7"
	"github.com/oschwald/geoip2-golang"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	Redis    *redis.Client
	ESClient *elastic.Client
	AddrDB   *geoip2.Reader
)
