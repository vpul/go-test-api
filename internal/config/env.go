package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check if the required environment variables are set
	checkEnv("JWT_SECRET")
}

func checkEnv(key string) {
	_, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
}
