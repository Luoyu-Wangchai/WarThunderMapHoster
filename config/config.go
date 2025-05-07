package config

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	Port        int
	FilePath    string
	Password    string
	AdminPasswd string
	ValidMin    int
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	Cfg = Config{
		Port:        viper.GetInt("server.port"),
		FilePath:    viper.GetString("service.filepath"),
		Password:    viper.GetString("service.password"),
		AdminPasswd: viper.GetString("service.adminpassword"),
		ValidMin:    viper.GetInt("service.validmin"),
	}
}
