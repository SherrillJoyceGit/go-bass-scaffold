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

func NewRedisClient() (*redis.Client, error) {

	cfg, err := getCurrentConfig()

	if err == nil {
		//log.Fatal(2, "Fail to get section 'database': %v", err)
		logger.CurrentLogger().Fatal("Fail to get section 'redis': " + err.Error())
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})

	_, connErr := client.Ping().Result()

	if connErr != nil {
		//panic("connect to redis failed！！！")
		logger.CurrentLogger().WithFields(logrus.Fields{
			"method": "redis-auth-connect",
		}).Panic("connect to " + cfg.Addr + " for redis failed！！！")
	} else {
		//fmt.Println("connect to redis OK!!!")
		logger.CurrentLogger().WithFields(logrus.Fields{
			"method": "redis-auth-connect",
		}).Infof("connect to " + cfg.Addr + " for redis is ok")
	}

	redisClient = client
	return client, connErr
}
