package middleware

import (
	"github.com/gin-gonic/gin"
	"goApi/global"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(GinRecovery(global.LOGGER, true))
	r.Use(Cors())
	r.Use(ParseParamMiddleware())
	r.Use(ApiLogMiddleware())
	sysConf := global.SERVER_CONFIG.System
	// nginx 不配置https 的话 可以在gin 自定义
	if sysConf.Ishttps {
		r.Use(LoadTls())
	}
}