package service

import (
	"errors"
	"github.com/BaiMeow/HduHelpLogin/models"
	"regexp"
)

var emailPattern = regexp.MustCompile("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")

func GetUserById(id uint) (*models.User, error) {
	user, err := models.GetUserById(id)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			return new(models.User), nil
		}
		return nil, err
	}
	return user, nil
}

func UpsertUser(user *models.User) error {
	if !emailPattern.MatchString(user.Email) || !(user.Age < 200) || !(user.Phone < 20000000000) || !(user.Phone > 10000000000) || user.ID == 0 {
		return ErrWrongFormat
	}
	return models.UpsertUser(user)
}
