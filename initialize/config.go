package initialize

import (
	"GinFrame/global"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {

	v := viper.New()

	v.SetConfigFile("./conf/settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		color.Red("[InitConfig]读取配置文件失败")
		panic(err)
	}

	serverConfig := global.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		color.Red("[InitConfig]配置文件解析失败")
		panic(err)
	}

	// 传递给全局变量
	global.GinConfig = serverConfig
	color.Blue("配置文件初始化完成", global.GinConfig.ServerInfo.LogPath)

}
