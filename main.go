package main

import (
	"github.com/BaiMeow/HduHelpLogin/conf"
	"github.com/BaiMeow/HduHelpLogin/models"
	"github.com/BaiMeow/HduHelpLogin/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {

}

func main() {
	conf.Init()
	models.Init()

	gin.SetMode(conf.Env.GetString("mode"))

	server := http.Server{
		Handler: router.InitRouter(),
		Addr:    ":" + conf.Env.GetString("server.port"),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
