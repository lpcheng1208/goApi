package global

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"github.com/zheng-ji/goSnowFlake"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"goApi/config"
)

var (
	REDIS_CONN    *redis.Client
	SERVER_CONFIG config.Server
	VIPER_CONFIG  *viper.Viper
	OPLOGGER      *oplogging.Logger
	LOGGER        *zap.Logger
	MysqlDb       *sql.DB
	MgoDB         *mgo.Session
	UniqId        *goSnowFlake.IdWorker
)

func LoggerWithContext(ctx *gin.Context) zap.Logger {
	newLogger := LOGGER
	if ctx != nil {
		RequestId := ctx.GetHeader("X-Request-ID")
		newLogger = newLogger.With(zap.String("RequestId", RequestId))
	}
	return *newLogger
}
