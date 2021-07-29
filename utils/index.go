package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string, defaultVal ...string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value := os.Getenv(key)
	if len(value) == 0 && len(defaultVal[0]) > 1 {
		value = defaultVal[0]
	}

	return value
}
