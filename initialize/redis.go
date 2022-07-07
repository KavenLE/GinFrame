package initialize

import (
	"GinFrame/global"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", global.GinConfig.RedisInfo.Host, global.GinConfig.RedisInfo.Port)
	// 生成redis客户端
	global.GinRedis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.GinConfig.RedisInfo.Password,
		DB:       13, // redis DB
	})
	// 链接redis
	_, err := global.GinRedis.Ping().Result()

	if err != nil {
		color.Red("[InitRedis] 链接redis异常:")
		panic(err)
	}
	color.Blue("redis链接成功")

}
