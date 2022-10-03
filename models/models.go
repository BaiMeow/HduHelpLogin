package models

import (
	"errors"
	"github.com/BaiMeow/HduHelpLogin/conf"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

var db *gorm.DB
var ErrDatabase = errors.New("database error,please contact server manager")
var ErrRecordNotFound = errors.New("record not found")

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open(conf.Env.GetString("database.source")), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to open database:%v", err)
	}
	if err := db.AutoMigrate(&Auth{}); err != nil {
		log.Fatalf("fail to migrate models:%v", err)
	}
	rand.Seed(time.Now().UnixNano())
}
