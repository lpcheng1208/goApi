package initialize

import (
	"gopkg.in/mgo.v2"
	"goApi/global"
	"os"
)

func InitMgoDb() {
	mongoConfig := global.SERVER_CONFIG.Mongo
	//MgoDb, err := mgo.Dial("mongodb://192.168.2.28:27017,192.168.2.28:27018,192.168.2.28:27019/?replicaSet=howie")
	mongoUrl := "mongodb://" + mongoConfig.UserName + ":" + mongoConfig.Password + "@" + mongoConfig.Addr + "/" + mongoConfig.DbName
	global.OPLOGGER.Debugf("mongoUrl :%s", mongoUrl)
	MgoDb, err := mgo.Dial(mongoUrl)
	if err != nil {
		global.OPLOGGER.Error("mgo", err)
		os.Exit(0)
	}
	global.MgoDB = MgoDb
	global.OPLOGGER.Debug("init mongo success")
}
