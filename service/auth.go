package service

import (
	"errors"
	"github.com/BaiMeow/HduHelpLogin/models"
	"github.com/google/uuid"
	"regexp"
)

var (
	UsernamePattern = regexp.MustCompile("^[a-zA-Z\\d_-]{4,16}$")
	PasswordPattern = regexp.MustCompile("^[a-zA-Z\\d_-]{4,16}$")
)

var (
	ErrUsernameExist = errors.New("用户名被占用")
	ErrWrongFormat   = errors.New("格式错误")
	ErrInvalidToken  = errors.New("无效的token")
)

var token = make(map[uuid.UUID]uint)

func Login(username, password string) (uint, error) {
	if !UsernamePattern.MatchString(username) || !PasswordPattern.MatchString(password) {
		return 0, ErrWrongFormat
	}
	return models.CheckAuth(username, password)
}

func Register(username, password string) (uint, error) {
	if !UsernamePattern.MatchString(username) || !PasswordPattern.MatchString(password) {
		return 0, ErrWrongFormat
	}
	id, err := models.AddAuth(username, password)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, ErrUsernameExist
	}
	return id, nil
}

func Logout(tk string) error {
	uu, err := uuid.Parse(tk)
	if err != nil {
		return ErrInvalidToken
	}

	delete(token, uu)
	return nil
}

// todo: token expire

func GetOrAddToken(id uint) string {
	for k, v := range token {
		if v == id {
			return k.String()
		}
	}
	uu := uuid.New()
	token[uu] = id
	return uu.String()
}

func GetIdByToken(tk string) (uint, error) {
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
