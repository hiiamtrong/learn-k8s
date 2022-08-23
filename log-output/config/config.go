package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Message string `mapstructure:"message" env:"MESSAGE"`
}

var Cfg Config

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Cfg)

	if err != nil {
		panic(err)
	}

	fmt.Println(Cfg)
}
