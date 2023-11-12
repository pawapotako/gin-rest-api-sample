package util

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
	DB  DatabaseConfig
}

type AppConfig struct {
	Port int
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func LoadConfig() Config {

	config := Config{}

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Unable to decode into struct", err)
	}

	return config
}
