package api

import (
	"errors"
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
	id := r.GetUint("id")
	user, err := service.GetUserById(id)
	if err != nil {
		r.String(http.StatusInternalServerError, "")
		return
	}
	r.JSON(http.StatusOK, &GetUserResp{
		user.Age, user.Phone, user.Email,
	})
}

func UpdateUser(r *gin.Context) {
	user := new(models.User)
	age, err := strconv.ParseUint(r.PostForm("age"), 10, 32)
	if err != nil {
		r.String(http.StatusBadRequest, "invalid input")
		return
	}
	user.Phone, err = strconv.ParseUint(r.PostForm("phone"), 10, 64)
	if err != nil {
		r.String(http.StatusBadRequest, "invalid input")
		return
	}
	user.Email = r.PostForm("email")
	user.Age = uint(age)
	if err := service.UpsertUser(user); err != nil {
		if errors.Is(err, service.ErrWrongFormat) {
			r.String(http.StatusBadRequest, "invalid input")
			return
		}
		r.String(http.StatusInternalServerError, "internal server error")
		return
	}
	r.String(http.StatusOK, "")
}
