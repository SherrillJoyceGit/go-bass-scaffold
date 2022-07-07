package db

import (
	"fmt"
	"github.com/SherrillJoyceGit/go-bass-scaffold/config"
	"github.com/SherrillJoyceGit/go-bass-scaffold/middle/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"log"
)

var Cloud *gorm.DB

func init() {
	var (
		err                                        error
		dbType, dbName, user, password, host, port string
	)

	//section, err := setting.Cfg.GetSection("database")

	if err != nil {
		//log.Fatal(2, "Fail to get section 'database': %v", err)
		logger.LoggerCurrent().Fatal("Fail to get section 'database': " + err.Error())
	}

	dbType = config.DbConfig.DbType
	dbName = config.DbConfig.DbName
	user = config.DbConfig.User
	password = config.DbConfig.Password
	host = config.DbConfig.Host
	port = config.DbConfig.Port
	ds := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	//jdbc:postgresql://10.10.76.215:5432/message?stringtype=unspecified
	//db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	Cloud, err = gorm.Open(dbType, ds)

	if err != nil {
		logger.LoggerCurrent().WithFields(logrus.Fields{
			"method": "cloud-connect",
		}).Panicln("connect to " + config.DbConfig.Host + " failed,err: " + err.Error())

		log.Panicln(err)
	}

	Cloud.SingularTable(true)
	Cloud.LogMode(true)
	Cloud.DB().SetMaxIdleConns(10)
	Cloud.DB().SetMaxOpenConns(100)

	logger.LoggerCurrent().WithFields(logrus.Fields{
		"method": "db-cloud-connect",
	}).Infof("connect to " + config.DbConfig.Host + " for " + config.DbConfig.DbName + " is ok")
}

func CloseCloudDB() {
	defer Cloud.Close()
}
