package initialize

import (
	"GinFrame/internal/router"
	"GinFrame/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 配置日志
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 配置跨域
	Router.Use(middlewares.Next())

	// 配置swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "服务启动成功！")
	})

	// 路由分组
	apiGroup := Router.Group("/v1/")
	router.InitBaseRouter(apiGroup)
	return Router
}
