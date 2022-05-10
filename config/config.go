package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig() Config {
	var cfg Config
	err := cleanenv.ReadConfig("application.yaml", &cfg)
	if err != nil {
		fmt.Println(err)
	}

	return cfg
}
