package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort    int
	Environment string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		if os.Getenv("ENVIRONMENT") != "production" {
			fmt.Println("Warning: .env file not found, using system env vars")
		}
	}

	config := &Config{
		HttpPort:    getEnvAsInt("HTTP_PORT", 8080),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) validate() error {
	// if _, err := strconv.Atoi(c.HttpPort); err != nil {
	// 	return fmt.Errorf("HTTP_PORT must be a valid port number, got %s", c.HttpPort)
	// }
	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}
