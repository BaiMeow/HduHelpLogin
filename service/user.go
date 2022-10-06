package service

import (
	"context"
	"errors"
	"github.com/BaiMeow/HduHelpLogin/models"
	"regexp"
)

var emailPattern = regexp.MustCompile("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")

var ErrWrongOldPassword = errors.New("旧密码错误")

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

func ChangePassword(ctx context.Context, id uint, oldPassword, newPassword string) error {
	ok, err := PasswordPattern.MatchString(oldPassword)
	if err != nil || !ok {
		return ErrWrongFormat
	}
	ok, err = models.CheckAuthWithId(ctx, id, oldPassword)
	if err != nil {
		return err
	}
	if !ok {
		return ErrWrongOldPassword
	}
	ok, err = PasswordPattern.MatchString(newPassword)
	if err != nil || !ok {
		return ErrWrongFormat
	}
	err = models.ChangeAuth(ctx, id, newPassword)
	if err != nil {
		return err
	}
	//logout
	token.Range(func(k, v any) bool {
		if v == id {
			token.Delete(k)
			return false
		}
		return true
	})
	return nil
}
