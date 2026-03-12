package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Port            string
	DatabaseURL     string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewConfig() *Config {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return &Config{
		Port:            os.Getenv("PORT"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 7 * 24 * time.Hour,
	}
}
