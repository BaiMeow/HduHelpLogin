package models

import (
	"errors"
	"github.com/BaiMeow/HduHelpLogin/utils"
	"gorm.io/gorm"
	"log"
)

type Auth struct {
	Username string
	Password []byte `gorm:"type:binary(8)"`
	Salt     []byte `gorm:"type:binary(20)"`
	gorm.Model
}

func CheckAuth(username, password string) (uint, error) {
	var auth Auth
	result := db.Where(Auth{Username: username}).First(&auth)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		log.Println(result.Error)
		return 0, ErrDatabase
	}
	if *(*[20]byte)(auth.Password) != utils.EncryptPassword(password, auth.Salt) {
		return 0, nil
	}
	return auth.ID, nil
}

func AddAuth(username, password string) (uint, error) {
	var auth Auth
	if err := db.Where(Auth{Username: username}).FirstOrInit(&auth).Error; err != nil {
		log.Println(err)
		return 0, ErrDatabase
	}
	if auth.ID != 0 {
		return 0, nil
	}
	auth.Salt = utils.GenSalt()
	p1 := utils.EncryptPassword(password, auth.Salt)
	auth.Password = p1[:]
	if err := db.Create(&auth).Error; err != nil {
		log.Println(err)
		return 0, ErrDatabase
	}
	return auth.ID, nil
}
