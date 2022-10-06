package models

import (
	"context"
	"errors"
	"github.com/BaiMeow/HduHelpLogin/utils"
	"gorm.io/gorm"
)

type Auth struct {
	Username string
	Password []byte `gorm:"type:binary(8)"`
	Salt     []byte `gorm:"type:binary(20)"`
	gorm.Model
}

func CheckAuth(ctx context.Context, username, password string) (uint, error) {
	var auth Auth
	result := db.WithContext(ctx).Where(Auth{Username: username}).First(&auth)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, ErrDatabase
	}
	if *(*[20]byte)(auth.Password) != utils.EncryptPassword(password, auth.Salt) {
		return 0, nil
	}
	return auth.ID, nil
}

func AddAuth(ctx context.Context, username, password string) (uint, error) {
	var auth Auth
	if err := db.WithContext(ctx).Where(Auth{Username: username}).FirstOrInit(&auth).Error; err != nil {
		return 0, ErrDatabase
	}
	if auth.ID != 0 {
		return 0, nil
	}
	auth.Salt = utils.GenSalt()
	p1 := utils.EncryptPassword(password, auth.Salt)
	auth.Password = p1[:]
	if err := db.WithContext(ctx).Create(&auth).Error; err != nil {
		return 0, ErrDatabase
	}
	return auth.ID, nil
}

func CheckAuthWithId(ctx context.Context, Id uint, password string) (bool, error) {
	var auth Auth
	result := db.WithContext(ctx).First(&auth, Id)
	if result.Error != nil {
		return false, ErrDatabase
	}
	if *(*[20]byte)(auth.Password) != utils.EncryptPassword(password, auth.Salt) {
		return false, nil
	}
	return true, nil
}

func ChangeAuth(ctx context.Context, Id uint, password string) error {
	var auth Auth
	result := db.WithContext(ctx).First(&auth, Id)
	if result.Error != nil {
		return ErrDatabase
	}
	auth.Salt = utils.GenSalt()
	p1 := utils.EncryptPassword(password, auth.Salt)
	auth.Password = p1[:]
	if err := db.WithContext(ctx).Save(&auth).Error; err != nil {
		return ErrDatabase
	}
	return nil
}
