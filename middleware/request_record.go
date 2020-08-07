package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goApi/config"
	"goApi/global"
	"time"
)

/**
api访问日志中间件
*/

func ApiLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		log := global.LoggerWithContext(ctx)
		ctx.Next()
		//请求耗时
		r := ctx.Request
		duration := time.Since(startTime)
		proto := r.Proto
		log.Info("用户访问日志",
			zap.String("ip", ctx.ClientIP()),
			zap.Any("method", r.Method),
			zap.Any("request path", r.URL.Path),
			zap.Any("proto", proto),
			zap.String("start", startTime.UTC().Format(config.TimeFormat)),
			zap.Duration("request cost", duration),
		)
	}

}
