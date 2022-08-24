package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PortClient string `mapstructure:"port_client" env:"PORT_CLIENT"`
	PortServer string `mapstructure:"port_server" env:"PORT_SERVER"`
	PGUser     string `mapstructure:"postgres_user" env:"POSTGRES_USER"`
	PGPassword string `mapstructure:"postgres_password" env:"POSTGRES_PASSWORD"`
	PGDB       string `mapstructure:"postgres_db" env:"POSTGRES_DB"`
	PGHost     string `mapstructure:"postgres_host" env:"POSTGRES_HOST"`
	PGPort     string `mapstructure:"postgres_port" env:"POSTGRES_PORT"`
	ApiKey     string `mapstructure:"api_key" env:"API_KEY"`
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

	fmt.Println(Cfg)
}
