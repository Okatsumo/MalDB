package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	URL                string `env:"APP_URL"`
	Port               int    `env:"APP_PORT"`
	Debug              bool   `env:"APP_DEBUG"`
	HttpRateLimit      int    `env:"HTTP_RATE_LIMIT"`
	HttpRateLimitBurst int    `env:"HTTP_RATE_LIMIT_BURST"`
}

// Load get env configs
func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Envivoment: .env file was not found in the project directory")
	}

	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Envivoment: Failed to load enviroment: %v", err)
	}

	return &cfg
}
