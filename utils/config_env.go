package utils

import (
	"fmt"
	"os"
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

func getEnv(key string, secret string) string {
	if secret != "" {
		data, err := os.ReadFile(secret)
		if err == nil {
			return string(data)
		}
	}
	return os.Getenv(key)
}

// Initialize common configurations from a .env file
func LoadConfig() (config CommonConfig) {
	config.MainServerAddr = getEnv("ADDR_ROUTE", "")
	config.MainServerPort = getEnv("PORT", "")
	config.ShortServerAddr = getEnv("ADDR_ROUTE_SHORTEN", "")
	config.ShortServerPort = getEnv("PORT_SHORTEN", "")
	config.MakeDBConfiguration()
	return
}

// Initialize database configurations into a string
func (conf *CommonConfig) MakeDBConfiguration() {
	dbSetup := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		getEnv("POSTGRES_HOST", ""),
		getEnv("POSTGRES_PORT", ""),
		getEnv("POSTGRES_USER", "/run/secrets/POSTGRES_USER"),
		getEnv("POSTGRES_NAME", "/run/secrets/POSTGRES_NAME"),
	)
	if pass := getEnv("POSTGRES_PASS", "/run/secrets/POSTGRES_PASS"); pass != "" {
		dbSetup = fmt.Sprintf("%s password=%s", dbSetup, pass)
	}
	conf.DBSourceName = dbSetup
}
