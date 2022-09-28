package router

import (
	"github.com/BaiMeow/HduHelpLogin/router/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	r.GET("/logout", api.Logout)
	//todo:getUser
	return r
}
