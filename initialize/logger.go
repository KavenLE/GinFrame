package initialize

import (
	"GinFrame/global"
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"time"
)

func InitLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%s/log_%s.log", global.GinConfig.ServerInfo.LogPath, GetNowFormatTodayTime()),
		"stdout",
	}
	logg, err := cfg.Build()
	if err != nil {
		color.Red("[InitLogger]初始化失败")
		panic(err)
	}
	zap.ReplaceGlobals(logg)
	color.Blue("日志管理初始化完成")
}

func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	return dateStr
}
