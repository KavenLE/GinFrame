package router

import (
	"GinFrame/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(r *gin.RouterGroup) {
	routerGroup := r.Group("base")
	{
		routerGroup.GET("test", controller.Base{}.Test)
	}
}
