package main

import (
	"goApi/core"
	"goApi/global"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {

	// 启动 数据库连接
	core.DbInit()
	// 程序结束前关闭数据库链接
	defer global.MysqlDb.Close()
	defer global.REDIS_CONN.Close()
	// 启动 http 后端服务程序
	core.RunHttpServer(1)

}
