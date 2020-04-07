package util

import (
	"log"
	"encoding/json"
	"os"
)

const (
	ConfigFileName         = "setting.json"
	OpenConfigFileErrMsg   = "打开配置文件失败"
	DecodeConfigFileErrMsg = "解析配置文件失败"
)

// 配置信息
type Config struct {
	DriverName           string
	DataSourceName       string
	RedisConnectAddress  string
	RedisConnectPassword string
}

// 解析配置文件
func ParseConfigFile() *Config {
	file, err := os.Open(ConfigFileName)
	if nil != err {
		log.Printf(OpenConfigFileErrMsg, err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if nil != err {
		log.Printf(DecodeConfigFileErrMsg, err)
	}
	return &config
}
