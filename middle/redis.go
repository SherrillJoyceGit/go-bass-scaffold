package middle

import (
	"github.com/SherrillJoyceGit/go-bass-scaffold/config"
	"github.com/SherrillJoyceGit/go-bass-scaffold/middle/logger"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func init() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AuthRedisConfig.Addr,
		Password: config.AuthRedisConfig.Password, // no password set
		DB:       config.AuthRedisConfig.Db,       // use default DB
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		//panic("connect to redis failed！！！")
		logger.LoggerCurrent().WithFields(logrus.Fields{
			"method": "redis-auth-connect",
		}).Panic("connect to " + config.AuthRedisConfig.Addr + " for redis failed！！！")
	} else {
		//fmt.Println("connect to redis OK!!!")
		logger.LoggerCurrent().WithFields(logrus.Fields{
			"method": "redis-auth-connect",
		}).Infof("connect to " + config.AuthRedisConfig.Addr + " for redis is ok")
	}
}
