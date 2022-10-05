package service

import (
	"context"
	"errors"
	"github.com/BaiMeow/HduHelpLogin/models"
	"regexp"
)

var emailPattern = regexp.MustCompile("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")

func GetUserById(ctx context.Context, id uint) (*models.User, error) {
	user, err := models.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			return new(models.User), nil
		}
		return nil, err
	}
	return user, nil
}

func UpsertUser(ctx context.Context, user *models.User) error {
	if !emailPattern.MatchString(user.Email) || !(user.Age < 200) || !(user.Phone < 20000000000) || !(user.Phone > 10000000000) || user.ID == 0 {
		return ErrWrongFormat
	}
	return models.UpsertUser(ctx, user)
}
