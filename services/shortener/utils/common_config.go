package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

type CommonConfig struct {
	ServerPort      string
	ServerAddr      string
	StaticShortURL  string
	DBHost          string
	DBPort          string
	DBName          string
	DBUser          string
	DBPass          string
	DBSourceName    string
	DBMigrationPath string
	DBSourceURL     string
	LogLevel        string
	AppEnv          string
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
func NewConfig() (config CommonConfig) {
	log.Info().Msg("Fetching common config")
	config.ServerAddr = getEnv("SERVER_ADDR", "")
	config.ServerPort = getEnv("SERVER_PORT", "")
	config.LogLevel = getEnv("LOG_LEVEL", "")
	config.StaticShortURL = getEnv("STATIC_SHORT_URL", "")
	config.AppEnv = getEnv("APP_ENV", "")
	config.DBHost = getEnv("POSTGRES_HOST", "")
	config.DBPort = getEnv("POSTGRES_PORT", "")
	config.DBUser = getEnv("POSTGRES_USER", "/run/secrets/POSTGRES_USER")
	config.DBName = getEnv("POSTGRES_NAME", "/run/secrets/POSTGRES_NAME")
	config.DBPass = getEnv("POSTGRES_PASS", "/run/secrets/POSTGRES_PASS")
	config.DBMigrationPath = getEnv("DB_MIGRATION_PATH", "")
	config.MakeDBConfiguration()
	config.MakeDBSourceURL()
	return
}

// Initialize database configurations into a string
func (conf *CommonConfig) MakeDBConfiguration() {
	dbSetup := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUser,
		conf.DBName,
	)
	if conf.DBPass != "" {
		dbSetup = fmt.Sprintf("%s password=%s", dbSetup, conf.DBPass)
	}
	conf.DBSourceName = dbSetup
}

func (conf *CommonConfig) MakeDBSourceURL() {
	conf.DBSourceURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.DBUser,
		conf.DBPass,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)
}
