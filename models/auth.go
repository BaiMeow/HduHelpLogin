package models

import (
	"github.com/BaiMeow/HduHelpLogin/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Username string
	Password []byte `gorm:"type:binary(8)"`
	Salt     []byte `gorm:"type:binary(20)"`
	gorm.Model
}

type ByteArray []byte

var ErrExistUsername = errors.New("username is already exist")
var ErrDatabase = errors.New("database error,please contact server manager")

func CheckAuth(username, password string) (bool, error) {
	var user User
	result := db.Where(User{Username: username}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		log.Println(result.Error)
		return false, ErrDatabase
	}
	if *(*[20]byte)(user.Password) == utils.EncryptPassword(password, user.Salt) {
		return true, nil
	}
	return false, nil
}

func AddAuth(username, password string) error {
	var user User
	if err := db.Where(User{Username: username}).FirstOrInit(&user).Error; err != nil {
		log.Println(err)
		return ErrDatabase
	}
	if user.ID != 0 {
		return ErrExistUsername
	}
	user.Salt = utils.GenSalt()
	p1 := utils.EncryptPassword(password, user.Salt)
	user.Password = p1[:]
	if err := db.Create(&user).Error; err != nil {
		log.Println(err)
		return ErrDatabase
	}
	return nil
}
