package router

import (
	"github.com/BaiMeow/HduHelpLogin/middlewave"
	"github.com/BaiMeow/HduHelpLogin/router/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	gin.Default()
	r.Use(middlewave.Logger(), gin.Recovery(), middlewave.AddTraceId)
	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	r.DELETE("/logout/:token", api.Logout)

	authed := r.Group("/api", middlewave.UserAuthentic)

	authed.GET("/user", api.GetUser)
	authed.PUT("/user", api.UpdateUser)
	authed.PUT("/user/password")

	return r
}
