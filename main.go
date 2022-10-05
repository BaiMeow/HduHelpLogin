package main

import (
	"github.com/BaiMeow/HduHelpLogin/conf"
	"github.com/BaiMeow/HduHelpLogin/log"
	"github.com/BaiMeow/HduHelpLogin/models"
	"github.com/BaiMeow/HduHelpLogin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Init()
	conf.Init()
	models.Init()

	gin.SetMode(conf.Env.GetString("mode"))

	err := router.InitRouter().Run(":" + conf.Env.GetString("server.port"))

	if err != nil {
		log.Logger.Fatalln(err)
	}

}
