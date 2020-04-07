package redis

import (
	"github.com/go-redis/redis"
	"grpc-sai-test/util"
	log "grpc-sai-test/store/log"

)

var RedisClient *redis.Client

const (
	RedisInitErrMsg    = "redis初始化失败"
	RedisInitFinishMsg = "redis初始化成功"
)

func init() {
	config := util.ParseConfigFile() // 解析配置文件
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisConnectAddress,
		Password: config.RedisConnectPassword,
	})
	res, err := client.Ping().Result() // ping
	if nil != err {
		log.Logger.Println(RedisInitErrMsg, err)
	} else {
		log.Logger.Println(RedisInitFinishMsg, res)
	}
	RedisClient = client
}
