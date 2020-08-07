package initialize

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"goApi/global"
)

func InitDbHelp() {
	mysqlConf := global.SERVER_CONFIG.Mysql
	Db, err := sql.Open("mysql", mysqlConf.Username+":"+mysqlConf.Password+"@("+mysqlConf.Path+")/"+mysqlConf.Dbname+"?"+mysqlConf.Config)
	if err != nil {
		panic(err.Error())
	}
	Db.SetMaxOpenConns(mysqlConf.MaxOpenConns)
	Db.SetMaxIdleConns(mysqlConf.MaxIdleConns)
	global.MysqlDb = Db
	global.LOGGER.Debug("init mysql success")
}
