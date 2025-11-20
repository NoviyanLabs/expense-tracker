package loader

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		Username     string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		SSLMode      string `mapstructure:"sslmode"`
		DatabaseName string `mapstructure:"name"`
	} `mapstructure:"database"`
	
	Application struct {
        ServerPort string `mapstructure:"port"`
    } `mapstructure:"app"`
}

func Load() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	v := viper.New()
	v.SetConfigFile("yaml")
	v.SetConfigName("config")
	v.AddConfigPath("./properties")

	if err := v.ReadInConfig(); err != nil {
		log.Print(err)
		panic("Error reading config")
	}

	v.SetConfigName("config-" + env)
	if err := v.MergeInConfig(); err != nil {
		log.Print("No config")
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatal("Error unmarshaling config")
	}

	return &cfg
}

func GetServerPort(config *Config) string {
	log.Print("Hello world")
	log.Print(config.Application.ServerPort)
	return config.Application.ServerPort
}
