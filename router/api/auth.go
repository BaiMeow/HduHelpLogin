package api

import (
	"fmt"
	"github.com/BaiMeow/HduHelpLogin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(r *gin.Context) {
	username := r.PostForm("username")
	password := r.PostForm("password")
	id, err := service.Login(username, password)
	if err != nil {
		r.String(http.StatusInternalServerError, fmt.Sprintf("internal server error:%v", err))
		return
	}
	if id == 0 {
		r.String(http.StatusUnauthorized, "用户名或密码错误")
		return
	}
	token := service.GetOrAddToken(id)
	if err != nil {
		r.String(http.StatusInternalServerError, fmt.Sprintf("internal server error:%v", err))
		return
	}
	r.JSON(http.StatusOK, struct {
		Token string `json:"token"`
	}{token})
}

func Register(r *gin.Context) {
	username := r.PostForm("username")
	password := r.PostForm("password")
	id, err := service.Register(username, password)
	if err != nil {
		r.String(http.StatusInternalServerError, fmt.Sprintf("internal server error:%v", err))
		return
	}
	if id == 0 {
		r.String(http.StatusForbidden, "注册失败:%v", err)
		return
	}
	r.String(http.StatusOK, "注册成功")
}

func Logout(r *gin.Context) {
	tk := r.Param("token")
	err := service.Logout(tk)
	if err != nil {
		r.String(http.StatusBadRequest, "invalid token")
		return
	}
	r.String(http.StatusOK, "logout success")
}
