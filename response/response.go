package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const defaultCode = 1001

type Body struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseCode(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Body{
		code, // 自定义code
		msg,  // message
		data, // 数据
	})
	c.Abort()
}

func Response(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Body{
		defaultCode, // 自定义code
		msg,         // message
		data,        // 数据
	})
	c.Abort()
}
