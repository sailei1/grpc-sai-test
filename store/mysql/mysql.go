package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"grpc-sai-test/util"
	log "grpc-sai-test/store/log"

)

var Engine *xorm.Engine

func init() {
	var err error
	config := util.ParseConfigFile()
	Engine, err = xorm.NewEngine(config.DriverName, config.DataSourceName)
	if nil != err {
		log.Logger.Println("mysql引擎初始化失败", err)
	} else {
		log.Logger.Println("mysql 初始化成功......")
	}
	Engine.ShowSQL(true)
}
