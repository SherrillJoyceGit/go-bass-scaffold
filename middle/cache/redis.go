package cache

import (
	"github.com/SherrillJoyceGit/go-bass-scaffold/middle/logger"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

var redisClient *redis.Client

func RedisClientCurrent() *redis.Client {
	return redisClient
}

func NewRedisClient(conf *RedisConfig) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password, // no password set
		DB:       conf.Db,       // use default DB
	})

	_, connErr := client.Ping().Result()

	if connErr != nil {
		//panic("connect to redis failed！！！")
		logger.LoggerCurrent().WithFields(logrus.Fields{
			"method": "redis-auth-connect",
		}).Panic("connect to " + conf.Addr + " for redis failed！！！")
	} else {
		//fmt.Println("connect to redis OK!!!")
		logger.LoggerCurrent().WithFields(logrus.Fields{
			"method": "redis-auth-connect",
		}).Infof("connect to " + conf.Addr + " for redis is ok")
	}

	redisClient = client
	return client, connErr
}
