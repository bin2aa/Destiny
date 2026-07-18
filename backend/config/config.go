package config

import (
	"os"
	"strconv"

	"destiny-backend/pkg/database"
)

// Config holds all application configuration.
type Config struct {
	Server   ServerConfig
	Postgres database.PostgresConfig
	Redis    *database.RedisConfig
	JWT      JWTConfig
}

// JWTConfig holds JWT configuration.
type JWTConfig struct {
	Secret string
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Host string
	Port int
}

// Load reads configuration from environment variables.
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvInt("SERVER_PORT", 8080),
		},
		Postgres: database.PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "123123"),
			DBName:   getEnv("DB_NAME", "Destiny"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Redis: &database.RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "destiny-secret-key-change-in-production"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
