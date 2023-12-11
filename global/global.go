package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"wild_goose_gin/config"
)

var (
	Config      *config.Config
	DB          *gorm.DB
	Logrus      *logrus.Logger
	ResMap      *config.ResMap
	RedisClient *redis.Client
	ESClient    *elastic.Client
)
