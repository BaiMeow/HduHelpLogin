package middlewave

import (
	"context"
	"errors"
	"github.com/BaiMeow/HduHelpLogin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UserAuthentic(r *gin.Context) {
	traceId := r.Value("traceId")
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	strs := strings.Fields(r.GetHeader("Authorization"))
	if len(strs) != 2 || strs[0] != "Bearer" {
		r.JSON(http.StatusUnauthorized, gin.H{
			"traceId": traceId,
			"msg":     "未登录",
		})
		r.Abort()
		return
	}
	id, err := service.GetIdByToken(ctx, strs[1])
	if err != nil {
		if errors.Is(err, service.ErrInvalidToken) {
			r.JSON(http.StatusUnauthorized, gin.H{
				"traceId": traceId,
				"msg":     "未登录",
			})
		} else {
			r.JSON(http.StatusInternalServerError, gin.H{
				"traceId": traceId,
				"msg":     "internal server error",
			})
		}
		r.Abort()
		return
	}
	r.Set("id", id)
}
