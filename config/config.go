package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Token string `mapstructure:"token"`
	DB    string `mapstructure:"database"`
}

func GetConfig() Config {
	viper.AddConfigPath("D:\\dev\\go\\go-line-notify\\config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	viper.Unmarshal(&config)

	return config
}
