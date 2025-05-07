package config

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	Port     int
	FilePath string
	Password string
	ValidMin int
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	Cfg = Config{
		Port:     viper.GetInt("server.port"),
		FilePath: viper.GetString("service.filepath"),
		Password: viper.GetString("service.password"),
		ValidMin: viper.GetInt("service.validmin"),
	}
}
