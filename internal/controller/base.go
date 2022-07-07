package controller

import (
	"GinFrame/response"
	"github.com/gin-gonic/gin"
)

type Base struct{}

func (b *Base) Test(c *gin.Context) {
	response.Response(c, "hello ginFrame", "[1,2,3]")
}
