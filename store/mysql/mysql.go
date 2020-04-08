package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "grpc-sai-test/store/log"
	"grpc-sai-test/util"
	"time"
)

var Engine *gorm.DB

func init() {
	var err error
	config := util.ParseConfigFile()
	Engine, err = gorm.Open(config.DriverName, config.DataSourceName)
	if nil != err {
		log.Logger.Println("mysql引擎初始化失败", err)
	} else {
		log.Logger.Println("mysql 初始化成功......")
	}
	Engine.DB().SetMaxIdleConns(5)
	Engine.DB().SetConnMaxLifetime(5 * time.Hour)
}
