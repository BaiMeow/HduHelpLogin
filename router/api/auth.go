package api

import (
	"context"
	"fmt"
	"github.com/BaiMeow/HduHelpLogin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(r *gin.Context) {
	traceId := r.Value("traceId")
	ctx := context.WithValue(context.Background(), "traceId", traceId)
	username := r.PostForm("username")
	password := r.PostForm("password")
	id, err := service.Login(ctx, username, password)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"traceId": traceId,
			"msg":     fmt.Sprintf("internal server error:%v", err),
		})
		return
	}
	if id == 0 {
		r.JSON(http.StatusUnauthorized, gin.H{
			"traceId": traceId,
			"msg":     "用户名或密码错误",
		})
		return
	}
	token := service.GetOrAddToken(ctx, id)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"traceId": traceId,
			"msg":     fmt.Sprintf("internal server error:%v", err),
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"traceId": traceId,
		"token":   token,
	})
}

func Register(r *gin.Context) {
	traceId := r.Value("traceId")
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	username := r.PostForm("username")
	password := r.PostForm("password")

	_, err := service.Register(ctx, username, password)
	if err != nil {
		r.JSON(http.StatusForbidden, gin.H{
			"traceId": traceId,
			"msg":     fmt.Sprintf("注册失败:%v", err),
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"traceId": traceId,
		"msg":     "注册成功",
	})
}

func Logout(r *gin.Context) {
	traceId := r.Value("traceId")
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	tk := r.Param("token")

	err := service.Logout(ctx, tk)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"traceId": traceId,
			"msg":     "invalid token",
		})
		return
	}
	r.JSON(http.StatusBadRequest, gin.H{
		"traceId": traceId,
		"msg":     "logout success",
	})
}
