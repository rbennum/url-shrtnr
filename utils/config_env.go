package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type CommonConfig struct {
	MainServerPort  string
	MainServerAddr  string
	ShortServerPort string
	ShortServerAddr string
	DBHost          string
	DBPort          string
	DBName          string
	DBUser          string
	DBPass          string
	DBSourceName    string
}

// Fetches the value of the environment key.
//
// This function provides a fallback in case the key doesn't exist.
func GetEnv(key string, fallback ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if len(fallback) == 0 {
		return ""
	}
	return fallback[0]
}

// Initialize common configurations from a .env file
func LoadConfig() (config CommonConfig, err error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		err = fmt.Errorf("expected the env to be set")
		return
	}
	if appEnv == "dev" {
		err = godotenv.Load("dev.env")
	} else {
		err = godotenv.Load("prod.env")
	}
	if err != nil {
		return
	}
	config.MainServerAddr = GetEnv("ADDR_ROUTE")
	config.MainServerPort = GetEnv("PORT")
	config.ShortServerAddr = GetEnv("ADDR_ROUTE_SHORTEN")
	config.ShortServerPort = GetEnv("PORT_SHORTEN")
	config.MakeDBConfiguration()
	return
}

// Initialize database configurations into a string
func (conf *CommonConfig) MakeDBConfiguration() {
	dbSetup := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		GetEnv("POSTGRES_HOST"),
		GetEnv("POSTGRES_PORT"),
		GetEnv("POSTGRES_USER"),
		GetEnv("POSTGRES_NAME"),
	)
	if pass := GetEnv("POSTGRES_PASS"); pass != "" {
		dbSetup = fmt.Sprintf("%s password=%s", dbSetup, pass)
	}
	conf.DBSourceName = dbSetup
}
