package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Fetches the value of the environment key.
//
// This function provides a fallback in case the key doesn't exist.
func GetEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

// Load any config in the .env file.
// For now we are putting our configs in the .env file
// unless stated otherwise.
//
// Call this function immediately before the app starts.
func LoadConfig(loc ...string) {
    err := godotenv.Load(loc...)
    if err != nil {
        log.Printf("Error loading .env file: %v", err)
    }
}