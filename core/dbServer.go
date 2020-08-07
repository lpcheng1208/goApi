package core

import (
	"goApi/initialize"
)

func DbInit() {
	// 启动原生sql 驱动
	initialize.InitDbHelp()
	// 初始化mgo db
	//initialize.InitMgoDb()
	// 初始化redis服务
	initialize.Redis()

}
