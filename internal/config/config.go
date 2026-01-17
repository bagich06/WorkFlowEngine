package config

import (
	"log"
	"os"
)

type Config struct {
	Port      string
	DBUrl     string
	JWTSecret string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	cfg := &Config{
		Port:      port,
		DBUrl:     os.Getenv("DATABASE_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
	if cfg.DBUrl == "" || cfg.JWTSecret == "" {
		log.Fatal("DATABASE_URL and JWT_SECRET must be set in environment")
	}
	return cfg
}
