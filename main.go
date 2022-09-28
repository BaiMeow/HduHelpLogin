package main

import (
	"github.com/BaiMeow/HduHelpLogin/conf"
	"github.com/BaiMeow/HduHelpLogin/models"
	"github.com/BaiMeow/HduHelpLogin/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

}

func main() {

	conf.Init()
	models.Init(conf.Env.GetString("database.source"))

	gin.SetMode(conf.Env.GetString("mode"))

	server := http.Server{
		Handler: router.InitRouter(),
		Addr:    ":" + conf.Env.GetString("server.port"),
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
