package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func GetString(key, defaultValue string) string {
	loadEnv()

	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}

func GetInt(key string, defaultValue int) int {
	loadEnv()
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}

	convertedValue, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}
	return convertedValue
}

func GetExpirationDuration() time.Duration {
	expirationStr := os.Getenv("TOKEN_EXPIRATION")
	duration, err := time.ParseDuration(expirationStr)
	if err != nil {
		log.Fatalf("Error while trying to parse duration value: %v", err)
	}
	return duration
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
