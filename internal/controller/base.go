package controller

import (
	"GinFrame/internal/forms"
	"GinFrame/response"
	"github.com/gin-gonic/gin"
)

type Base struct{}

// Test
// @Summary 介绍
// @Accept json
// @Tags Base
// @Security Bearer
// @Produce  form
// @Param Param path int true "param"
// @Resource Name
// @Router /base/test [get]
// @Success 1001
func (b *Base) Test(c *gin.Context) {
	form := forms.TestForm{}
	if form.Param == "" {
		response.Response(c, "参数错误", "")
		return
	}
	response.Response(c, "hello ginFrame", form.Param)
	return
}
