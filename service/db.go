package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"salotto/model"
	"time"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "test_h.db")
	if err = DB.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(50)
	DB.DB().SetConnMaxLifetime(5 * time.Minute)
}
