package service

import (
	"context"
	"errors"
	"github.com/BaiMeow/HduHelpLogin/models"
	"github.com/dlclark/regexp2"
	"github.com/google/uuid"
	"regexp"
)

var (
	UsernamePattern = regexp.MustCompile(`^[a-zA-Z\d_-]{4,16}$`)
	// PasswordPattern 密码正则，最少6位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符
	PasswordPattern = regexp2.MustCompile(`^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*? ]).*$`, 0)
)

var (
	ErrUsernameExist = errors.New("用户名被占用")
	ErrWrongFormat   = errors.New("格式错误")
	ErrInvalidToken  = errors.New("无效的token")
)

var token = make(map[uuid.UUID]uint)

func Login(ctx context.Context, username, password string) (uint, error) {
	if !UsernamePattern.MatchString(username) {
		return 0, ErrWrongFormat
	}
	if ok, _ := PasswordPattern.MatchString(password); !ok {
		return 0, ErrWrongFormat
	}
	return models.CheckAuth(ctx, username, password)
}

func Register(ctx context.Context, username, password string) (uint, error) {
	if !UsernamePattern.MatchString(username) {
		return 0, ErrWrongFormat
	}
	if ok, _ := PasswordPattern.MatchString(password); !ok {
		return 0, ErrWrongFormat
	}
	id, err := models.AddAuth(ctx, username, password)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, ErrUsernameExist
	}
	return id, nil
}

func Logout(ctx context.Context, tk string) error {
	uu, err := uuid.Parse(tk)
	if err != nil {
		return ErrInvalidToken
	}
	delete(token, uu)
	return nil
}

// todo: token expire

func GetOrAddToken(ctx context.Context, id uint) string {
	for k, v := range token {
		if v == id {
			return k.String()
		}
	}
	uu := uuid.New()
	token[uu] = id
	return uu.String()
}

func GetIdByToken(ctx context.Context, tk string) (uint, error) {
	uu, err := uuid.Parse(tk)
	if err != nil {
		return 0, ErrInvalidToken
	}
	id, ok := token[uu]
	if !ok {
		return 0, ErrInvalidToken
	}
	return id, nil
}
