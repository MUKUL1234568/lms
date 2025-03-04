package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func LoadConfig() {
	err := godotenv.Load("/home/xs525-mukcha/Desktop/lms/api/.env") // Adjust path if needed
	if err != nil {
		log.Println("⚠️  No .env file found. Using default settings.")
	} else {
		log.Println("✅ .env file loaded successfully.")
	}
}

// GetEnv gets the environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
