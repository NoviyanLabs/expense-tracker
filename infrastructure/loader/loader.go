package loader

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
}

func Load() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = ""
	}

	v := viper.New()
	v.SetConfigFile("yaml")
	v.SetConfigFile("config")
	v.AddConfigPath("../../properties")

	if err := v.ReadInConfig(); err!= nil {
		panic("Error reading config")
	}

	
}