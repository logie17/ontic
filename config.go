package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Config = viper.New()

func ReadConfig() {
	Config.SetConfigName("conf")
	Config.AddConfigPath(".")
	err := Config.ReadInConfig()
	if err != nil {
		log.Panic(fmt.Errorf("Unable to read conf file %v", err))
	}
}
