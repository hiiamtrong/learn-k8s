package config

import "github.com/spf13/viper"

type Config struct {
	PortClient string `mapstructure:"port_client"`
	PortServer string `mapstructure:"port_server"`
}

var Cfg Config

func Init() {
	viper.SetDefault("port_client", "8080")
	viper.SetDefault("port_server", "9090")
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
