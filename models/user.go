package models

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	Age   uint
	Phone uint64
	Email string
	gorm.Model
}

func GetUserById(ctx context.Context, id uint) (*User, error) {
	var user = new(User)
	user.ID = id
	result := db.WithContext(ctx).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, ErrDatabase
	}
	return user, nil
}

func UpsertUser(ctx context.Context, user *User) error {
	result := db.WithContext(ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(user)
	if result.Error != nil {
		return ErrDatabase
	}
	return nil
}
