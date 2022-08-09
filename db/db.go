package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgresDb(cfg *Config) *gorm.DB {
	var (
		dbName, userName, password, host, port string
	)

	dbName = cfg.DbName
	userName = cfg.UserName
	password = cfg.Password
	host = cfg.Host
	port = cfg.Port
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, userName, dbName, password)
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Printf("connect to " + cfg.Host + " failed,err: " + err.Error())
		log.Panicln(err)
		return nil
	} else {
		//db.DB().SetMaxIdleConns(10)
		//db.DB().SetMaxOpenConns(100)
		log.Printf("connecting to " + cfg.Host + " for " + cfg.DbName + " is successful")
		return db
	}

}
