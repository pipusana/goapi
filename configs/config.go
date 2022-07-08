package configs

import (
	"log"

	"github.com/pipusana/goapi/entities"
	"github.com/spf13/viper"
)

func ReadConfig() entities.Config {
	log.Println("Loading Server Configurations...")
	var appConfig entities.Config
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatal(err)
	}

	return appConfig
}
