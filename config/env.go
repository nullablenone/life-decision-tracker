package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	DBSSLMode string

	PortLocal string
}

func NewEnv() (*Env, error) {
	err := godotenv.Load(".env")

	app_env := os.Getenv("APP")
	if app_env == "" {
		app_env = "local"
	}

	if err != nil && app_env == "local" {
		return nil, fmt.Errorf("failed to load .env file in local: %w", err)
	}

	env := Env{
		DBHost:    os.Getenv("DB_HOST"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		DBPort:    os.Getenv("DB_PORT"),
		DBSSLMode: os.Getenv("DB_SSLMODE"),

		PortLocal: os.Getenv("PORT_Local"),
	}

	return &env, nil
}
