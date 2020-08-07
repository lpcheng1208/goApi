package response

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goApi/global"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	NOTLOGIN = 13
	ERROR    = 7
	SUCCESS  = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	// 获取带 Context 的 logger
	log := global.LoggerWithContext(c)
	log.Info("requestResult", zap.Any("request path", c.Request.URL.Path), zap.Int("code", code), zap.String("msg", msg))
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
