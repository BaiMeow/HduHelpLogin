package middlewave

import (
	"errors"
	"github.com/BaiMeow/HduHelpLogin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UserAuthentic(r *gin.Context) {
	strs := strings.Fields(r.GetHeader("Authorization"))
	if len(strs) != 2 || strs[0] != "Bearer" {
		r.String(http.StatusUnauthorized, "未登录")
		r.Done()
		return
	}
	id, err := service.GetIdByToken(strs[1])
	if err != nil {
		if errors.Is(err, service.ErrInvalidToken) {
			r.String(http.StatusUnauthorized, "未登录")
		} else {
			r.String(http.StatusInternalServerError, "internal server fail")
		}
		r.Abort()
		return
	}
	r.Set("id", id)
}
