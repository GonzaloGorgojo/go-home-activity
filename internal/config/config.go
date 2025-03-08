package config

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecretKey string
	JWTIssuer    string
	DBConnection string
	Port         string
}

var (
	once     sync.Once
	instance *Config
)

func LoadConfig() (*Config, error) {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Warning: No .env file found, using system environment variables")
		}

		instance = &Config{
			JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
			DBConnection: os.Getenv("DATABASE_OPEN"),
			JWTIssuer:    os.Getenv("JWT_ISSUER"),
			Port:         os.Getenv("PORT"),
		}
	})

	if instance.JWTSecretKey == "" {
		return nil, errors.New("ERROR: JWT_SECRET_KEY is required but missing")
	}
	if instance.DBConnection == "" {
		return nil, errors.New("ERROR: DATABASE_OPEN is required but missing")
	}
	if instance.JWTIssuer == "" {
		return nil, errors.New("ERROR: JWT_ISSUER is required but missing")
	}
	if instance.Port == "" {
		return nil, errors.New("ERROR: PORT is required but missing")
	}

	return instance, nil
}
