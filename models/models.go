package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

var db *gorm.DB

func Init(source string) {
	var err error
	db, err = gorm.Open(sqlite.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to open database:%v", err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("fail to migrate models:%v", err)
	}
	rand.Seed(time.Now().UnixNano())
}
