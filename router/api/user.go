package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/BaiMeow/HduHelpLogin/models"
	"github.com/BaiMeow/HduHelpLogin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetUserResp struct {
	Age   uint   `json:"age"`
	Phone uint64 `json:"phone"`
	Email string `json:"email"`
}

func GetUser(r *gin.Context) {
	traceId := r.Value("traceId")
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	id := r.GetUint("id")

	user, err := service.GetUserById(ctx, id)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"traceId": traceId,
			"msg":     fmt.Sprintf("internal server error:%v", err),
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"traceId": traceId,
		"msg":     "ok",
		"age":     user.Age,
		"phone":   user.Phone,
		"email":   user.Email,
	})
}

func UpdateUser(r *gin.Context) {
	traceId := r.Value("traceId")
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	user := new(models.User)
	age, err := strconv.ParseUint(r.PostForm("age"), 10, 32)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"traceId": traceId,
			"msg":     "invalid input",
		})
		return
	}
	user.Phone, err = strconv.ParseUint(r.PostForm("phone"), 10, 64)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"traceId": traceId,
			"msg":     "invalid input",
		})
		return
	}
	user.Email = r.PostForm("email")
	user.Age = uint(age)
	user.ID = r.GetUint("id")
	if err := service.UpsertUser(ctx, user); err != nil {
		if errors.Is(err, service.ErrWrongFormat) {
			r.JSON(http.StatusBadRequest, gin.H{
				"traceId": traceId,
				"msg":     "invalid input",
			})
			return
		}
		r.JSON(http.StatusInternalServerError, gin.H{
			"traceId": traceId,
			"msg":     "internal server error",
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"traceId": traceId,
		"msg":     "put success",
	})
}
