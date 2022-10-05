package models

import (
	"errors"
	"github.com/BaiMeow/HduHelpLogin/conf"
	"github.com/BaiMeow/HduHelpLogin/log"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

var db *gorm.DB
var ErrDatabase = errors.New("database error,please contact server manager")
var ErrRecordNotFound = errors.New("record not found")

func Init() {
	var err error
	lv := conf.Env.GetString("log.level")
	var level log.LogLevel
	switch lv {
	case "Silent":
		level = log.Silent
	case "Warn":
		level = log.Warn
	case "Error":
		level = log.Error
	case "Info":
		level = log.Info
	default:
		level = log.Warn
	}
	db, err = gorm.Open(sqlite.Open(conf.Env.GetString("database.source")), &gorm.Config{
		Logger: log.NewGormLogger(level),
	})
	if err != nil {
		logrus.Fatalf("fail to open database:%v", err)
	}
	if err := db.AutoMigrate(&Auth{}, &User{}); err != nil {
		logrus.Fatalf("fail to migrate models:%v", err)
	}
	rand.Seed(time.Now().UnixNano())
}
