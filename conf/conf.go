package conf

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Env *viper.Viper

//go:embed default_config.yaml
var defaultConf []byte

func Init() {
	Env = viper.New()
	Env.SetConfigType("yaml")
	Env.SetConfigName("config")
	Env.AddConfigPath(".")
	if err := Env.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := os.WriteFile("./config.yaml", defaultConf, 0666); err != nil {
				log.Fatalf("fail to write config:%v", err)
			}
			_ = Env.ReadConfig(bytes.NewReader(defaultConf))
		} else {
			log.Fatalf("fail to read config:%v", err)
		}
	}
}
