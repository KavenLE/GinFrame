package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	GinConfig ServerConfig
	GinDB     *gorm.DB
	GinRedis  *redis.Client
	GinLog    *zap.Logger
	// 全局翻译实例
	Trans ut.Translator
)
