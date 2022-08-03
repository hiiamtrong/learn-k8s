package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"port"`
}

var Cfg Config

func Init() {
	viper.SetDefault("port", "8080")
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
}
