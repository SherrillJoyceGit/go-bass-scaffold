package db

import (
	"fmt"
	"github.com/SherrillJoyceGit/go-bass-scaffold/middle/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"log"
)

var db *gorm.DB

func NewDbAccess() (*gorm.DB, error) {
	var (
		err                                        error
		dbType, dbName, user, password, host, port string
	)

	//section, err := setting.Cfg.GetSection("database")
	cfg, err := getCurrentConfig()
	if err == nil {
		//log.Fatal(2, "Fail to get section 'database': %v", err)
		logger.LoggerCurrent().Fatal("Fail to get section 'database': " + err.Error())
	}

	dbType = cfg.DbType
	dbName = cfg.DbName
	user = cfg.Username
	password = cfg.Password
	host = cfg.Host
	port = cfg.Port
	ds := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	//jdbc:postgresql://10.10.76.215:5432/message?stringtype=unspecified
	//db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	db, err = gorm.Open(dbType, ds)

	if err != nil {
		logger.LoggerCurrent().WithFields(logrus.Fields{
			"method": "cloud-connect",
		}).Panicln("connect to " + cfg.Host + " failed,err: " + err.Error())

		log.Panicln(err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	logger.LoggerCurrent().WithFields(logrus.Fields{
		"method": "db-cloud-connect",
	}).Infof("connect to " + cfg.Host + " for " + cfg.DbName + " is ok")

	return db, nil
}

func CloseCloudDB() {
	defer db.Close()
}
