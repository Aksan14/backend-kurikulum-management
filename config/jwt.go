package config

import (
	"os"
	"time"
)

type JWTConfig struct {
	Secret         string
	AccessExpires  time.Duration
	RefreshExpires time.Duration
	Issuer         string
}

var JWT *JWTConfig

func LoadJWTConfig() *JWTConfig {
	accessExp, _ := time.ParseDuration(getEnvOrDefault("JWT_EXPIRES_IN", "24h"))
	refreshExp, _ := time.ParseDuration(getEnvOrDefault("JWT_REFRESH_EXPIRES_IN", "168h"))

	JWT = &JWTConfig{
		Secret:         getEnvOrDefault("JWT_SECRET", "your-super-secret-jwt-key"),
		AccessExpires:  accessExp,
		RefreshExpires: refreshExp,
		Issuer:         getEnvOrDefault("APP_NAME", "Kurikulum Management System"),
	}
	return JWT
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}