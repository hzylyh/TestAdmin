package cronUtil

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"salotto/model"
	"time"
)

var CronDB *gorm.DB

func init() {
	var err error

	CronDB, err = gorm.Open("sqlite3", "test_h.db")
	if nil != err {
		log.Fatalf("opens database failed: " + err.Error())
	}

	if err = CronDB.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}

	CronDB.DB().SetMaxIdleConns(10)
	CronDB.DB().SetMaxOpenConns(50)
	CronDB.DB().SetConnMaxLifetime(5 * time.Minute)
	CronDB.LogMode(true)
}
