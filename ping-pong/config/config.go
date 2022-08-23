package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PGUser     string `mapstructure:"postgres_user" env:"POSTGRES_USER"`
	PGPassword string `mapstructure:"postgres_password" env:"POSTGRES_PASSWORD"`
	PGDB       string `mapstructure:"postgres_db" env:"POSTGRES_DB"`
	PGHost     string `mapstructure:"postgres_host" env:"POSTGRES_HOST"`
	PGPort     string `mapstructure:"postgres_port" env:"POSTGRES_PORT"`
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
