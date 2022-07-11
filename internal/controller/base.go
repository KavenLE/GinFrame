package controller

import (
	"GinFrame/internal/forms"
	"GinFrame/middlewares"
	"GinFrame/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Base struct{}

// Test
// @Summary 介绍
// @Accept json
// @Tags Base
// @Security Bearer
// @Produce  json
// @Param Param path int true "param"
// @Resource Name
// @Router /base/test [get]
// @Success 1001
func (b *Base) Test(c *gin.Context) {
	form := forms.TestForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		middlewares.HandleValidatorError(c, err)
		return
	}
	if form.Param == "" {
		response.Response(c, "参数错误", "")
		return
	}
	zap.L().Debug("this is test log", zap.String("user", "test"), zap.Int("age", 18))
	response.Response(c, "hello ginFrame", form.Param)
	return
}
