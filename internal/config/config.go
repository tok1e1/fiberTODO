package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type ServerConfig struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port int    `yaml:"port" env-default:"8080"`
}

type DatabaseConfig struct {
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true" env:"DB_PASSWORD"`
	Name     string `yaml:"name" env-required:"true"`
	Type     string `yaml:"type" env-default:"postgres"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("config file does not exist: %s", configPath)
		} else {
			log.Fatalf("error checking config file: %s", err)
		}
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
