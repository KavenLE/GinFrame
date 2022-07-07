package main

import (
	"GinFrame/global"
	"GinFrame/initialize"
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
)

// @title swagger
// @version 1.0
// @description browser server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8000
// @BasePath /v1

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitMysqlDB()
	initialize.InitRedis()
	Router := initialize.Routers()
	color.Cyan("=======================>服务启动成功<=======================")
	err := Router.Run(fmt.Sprintf(":%d", global.GinConfig.ServerInfo.Port))
	if err != nil {
		zap.L().Info("program run error!", zap.String("error", "启动错误!"))
	}
}
