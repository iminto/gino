package model

import (
	"gindemo/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)
var Db *gorm.DB
func InitDB(config *config.WebConfig) {
	var db *gorm.DB
	db,err:=gorm.Open(config.DbType,config.DSN)
	if err != nil {
		panic(err)
	}
	if config.SQLDebug {
		db.LogMode(true)
		db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
	Db=db
}
