package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Dot map[string]string
type Dots []Dot

var Config Dots

func ReadConfig() {
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(fmt.Errorf("Unable to read conf file %v", err))

	}

	if err := viper.UnmarshalKey("dots", &Config); err != nil {
		log.Panic(fmt.Errorf("Unable to unmarshal conf file %v", err))
	}
}
