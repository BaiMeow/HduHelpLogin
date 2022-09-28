package models

import (
	"github.com/BaiMeow/HduHelpLogin/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Username string
	Password [20]byte `gorm:"type:binary(20)"`
	Salt     [8]byte  `gorm:"type:binary(8)"`
	gorm.Model
}

func CheckUser(username, password string) (bool, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		log.Println(result.Error)
		return false, errors.New("database error,please contact server manager")
	}
	if user.Password == utils.EncryptPassword(password, user.Salt) {
		return true, nil
	}
	return false, nil
}
