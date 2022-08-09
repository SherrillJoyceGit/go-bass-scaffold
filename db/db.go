package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func NewDbAccess(cfg *Config) *gorm.DB {
	var (
		err                                            error
		dbType, dbName, userName, password, host, port string
	)

	dbType = cfg.DbType
	dbName = cfg.DbName
	userName = cfg.UserName
	password = cfg.Password
	host = cfg.Host
	port = cfg.Port
	ds := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, userName, dbName, password)
	db, err := gorm.Open(dbType, ds)

	if err != nil {
		/*		logger.CurrentLogger().WithFields(logrus.Fields{
				"method": "cloud-connect",
			}).Panicln("connect to " + cfg.Host + " failed,err: " + err.Error())*/
		log.Printf("connect to " + cfg.Host + " failed,err: " + err.Error())
		log.Panicln(err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	/*	logger.CurrentLogger().WithFields(logrus.Fields{
		"method": "db-cloud-connect",
	}).Infof("connect to " + cfg.Host + " for " + cfg.DbName + " is ok")*/
	log.Printf("connecting to " + cfg.Host + " for " + cfg.DbName + " is successful")
	return db
}

func CloseCloudDB(db *gorm.DB) {
	defer db.Close()
}
