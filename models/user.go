package models

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type User struct {
	Age   uint
	Phone uint64
	Email string
	gorm.Model
}

func GetUserById(id uint) (*User, error) {
	var user = new(User)
	user.ID = id
	result := db.First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		log.Println(result.Error)
		return nil, ErrDatabase
	}
	return user, nil
}

func UpsertUser(user *User) error {
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(user)
	if result.Error != nil {
		log.Println(result.Error)
		return ErrDatabase
	}
	return nil
}
