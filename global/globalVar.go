package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	GinConfig ServerConfig
	GinDB     *gorm.DB
	GinRedis  *redis.Client
	GinLog    *zap.Logger
)
