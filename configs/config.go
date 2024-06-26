package configs

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// JWT token configuration
	JWTConfig struct {
		ApiSecret   string `envconfig:"API_SECRET"`
		ExpireDelta int    `envconfig:"EXPIRE_DELTA"`
	}

	// Application configuration
	AppConfig struct {
		Port   int    `envconfig:"VERISAFE_PORT"`
		Addres string `envconfig:"VERISAFE_ADDRESS"`
	}

	// Database configuration
	DatabaseConfig struct {
		DatabaseHost     string `envconfig:"DB_HOST"`
		DatabaseDriver   string `envconfig:"DB_DRIVER"`
		DatabaseUser     string `envconfig:"DB_USER"`
		DatabasePassword string `envconfig:"DB_PASSWORD"`
		DatabaseName     string `envconfig:"DB_NAME"`
		DatabasePort     int    `envconfig:"DB_PORT"`
	}
}

// The LoadConfig function loads the env file specified and returns
// a valid configuration object ready for use
func LoadConfig() (*Config, error) {
	cfg := Config{}

	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("Failed to load environment variables: %v", err)
	}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("Failed to load environment variables: %v", err)
	}

	return &cfg, nil
}
