package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	URL                string `env:"APP_URL"`
	Debug              bool   `env:"APP_DEBUG"`
	HttpRateLimit      int    `env:"HTTP_RATE_LIMIT"`
	HttpRateLimitBurst int    `env:"HTTP_RATE_LIMIT_BURST"`
}

// Load get env configs
func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load enviroment: %v", err)
	}

	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Failed to load enviroment: %v", err)
	}

	return &cfg
}
