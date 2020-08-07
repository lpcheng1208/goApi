package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"go.uber.org/zap"
	"goApi/global"
	"goApi/initialize"
	"time"
)

func RunHttpServer(serverId int64) {

	// 初始化雪花id
	initialize.InitIw(serverId)

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.SERVER_CONFIG.System.Addr)

	s := endless.NewServer(address, Router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOGGER.Sugar().Debugf("server run success on %s", address)
	env := global.VIPER_CONFIG.Get("system.env")
	global.LOGGER.Debug("当前运行环境为", zap.Any("env", env))
	global.LOGGER.Sugar().Error(s.ListenAndServe())

}
