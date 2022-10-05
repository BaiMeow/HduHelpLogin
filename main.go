package main

import (
	"github.com/BaiMeow/HduHelpLogin/conf"
	"github.com/BaiMeow/HduHelpLogin/log"
	"github.com/BaiMeow/HduHelpLogin/models"
	"github.com/BaiMeow/HduHelpLogin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	models.Init()
	log.Init()

	gin.SetMode(conf.Env.GetString("mode"))

	err := router.InitRouter().Run(":" + conf.Env.GetString("server.port"))

	if err != nil {
		log.Logger.Fatalln(err)
	}

}
